// Initialize Firebase
var config = {
  apiKey: "xxx",
  authDomain: "xxx",
  databaseURL: "xxx",
  projectId: "sparring-hasura",
  storageBucket: "xxx",
  messagingSenderId: "xxx",
  appId: "xxx"
};
const provider = new firebase.auth.GoogleAuthProvider();

firebase.initializeApp(config);

document.getElementById('login-form').onsubmit = function(event) {
  event.preventDefault();
  login();
};


document.getElementById('get-token').onclick = function(event) {
  event.preventDefault();
  firebase.auth().currentUser.getIdToken(true).
    then(token => document.getElementById('id-token').innerHTML = token);
};

function login() {
  firebase.auth().signInWithPopup(provider)
    .then(function(user) {
      console.log('login success');
    })
    .catch(function(error) {
      // Handle Errors here.
      var errorCode = error.code;
      var errorMessage = error.message;
      console.log(error);
    });

  let callback = null;
  let metadataRef = null;
  firebase.auth().onAuthStateChanged(user => {
    // Remove previous listener.
    if (callback) {
      metadataRef.off('value', callback);
    }
    // On user login add new listener.
    if (user) {
      // Check if refresh is required.
      metadataRef = firebase.database().ref('metadata/' + user.uid + '/refreshTime');
      callback = (snapshot) => {
        // Force refresh to pick up the latest custom claims changes.
        // Note this is always triggered on first call. Further optimization could be
        // added to avoid the initial trigger when the token is issued and already contains
        // the latest claims.
        user.getIdToken(true);
      };
      // Subscribe new listener to changes on that node.
      metadataRef.on('value', callback);
    }
  });
}
