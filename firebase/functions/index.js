const functions = require("firebase-functions");
const admin = require("firebase-admin");
const axios = require("axios");

admin.initializeApp(functions.config().firebase);

const createUser = `
mutation createUser($id: String = "", $name: String = "", $email: String = "") {
  insert_users_one(object: { id: $id, name: $name, email: $email }, on_conflict: {constraint: users_pkey, update_columns: []}) {
    id
    email
    name
  }
}
`

const HASURA_GRAPHQL_ENDPOINT = functions.config().hasura.url
const HASURA_GRAPHQL_ADMIN_SECRET = functions.config().hasura.admin_secret

// On sign up.
exports.processSignUp = functions.auth.user().onCreate(user => {
  console.dir(`id:${user.uid} name:${user.name} email:${user.email}`)

  const customClaims = {
    "https://hasura.io/jwt/claims": {
      "x-hasura-default-role": "user",
      "x-hasura-allowed-roles": ["user"],
      "x-hasura-user-id": user.uid
    }
  };

  return admin
    .auth()
    .setCustomUserClaims(user.uid, customClaims)
    .then(() => {
      let query = {
        query: createUser,
        variables: {
          id: user.uid,
          name: user.name,
          email: user.email,
        },
      }
      axios({
        method: 'post',
        url: HASURA_GRAPHQL_ENDPOINT,
        data: query,
        headers: {
          'x-hasura-admin-secret': HASURA_GRAPHQL_ADMIN_SECRET,
        },
      })

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
