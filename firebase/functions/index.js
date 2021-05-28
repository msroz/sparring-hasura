const functions = require("firebase-functions");
const admin = require("firebase-admin");
const axios = require("axios");

admin.initializeApp(functions.config().firebase);

const MutationRegisterUser = `
  mutation RegisterUser($email: String!, $uid: String!) {
    registerUser(uid: $uid, email: $email) {
      user_id
    }
  }
`

const HASURA_GRAPHQL_ENDPOINT = functions.config().hasura.url
const HASURA_GRAPHQL_ADMIN_SECRET = functions.config().hasura.admin_secret

// On sign up.
exports.processSignUp = functions.auth.user().onCreate(async user => {
  console.log(`id:${user.uid} name:${user.displayName} email:${user.email}`)


  let query = {
    query: MutationRegisterUser,
    variables: {
      uid: user.uid,
      email: user.email,
    },
  }
  response = await axios({
    method: 'post',
    url: HASURA_GRAPHQL_ENDPOINT,
    data: query,
    headers: {
      'x-hasura-admin-secret': HASURA_GRAPHQL_ADMIN_SECRET,
    },
  })

  console.log(response.data)
  // { errors: [ { extensions: [Object], message: 'internal error' } ] }
  // { data: { registerUser: { user_id: '01F6RW0AH0PJ4EXQH1CNP16PRM' } } }

  const body = response.data
  if (body.errors) {
    // TODO: error
  }
  hasuraUserId = body.data.registerUser.user_id

  const customClaims = {
    "https://hasura.io/jwt/claims": {
      "x-hasura-default-role": "user",
      "x-hasura-allowed-roles": ["user"],
      "x-hasura-user-id": hasuraUserId,
    }
  };

  return admin
    .auth()
    .setCustomUserClaims(user.uid, customClaims)
    .then(() => {

      // Update real-time database to notify client to force refresh.
      const metadataRef = admin.database().ref("metadata/" + user.uid);
      // Set the refresh time to the current UTC timestamp.
      // This will be captured on the client to force a token refresh.
      return metadataRef.set({ refreshTime: new Date().getTime() });
    })
    .catch(error => {
      console.log(error);
    });
});
