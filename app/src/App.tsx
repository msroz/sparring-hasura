import React, { useState, useEffect } from 'react'
import jwt from 'jsonwebtoken'
import firebase from 'firebase/app'
import 'firebase/auth'
import 'firebase/database'

const provider = new firebase.auth.GoogleAuthProvider()

const firebaseConfig = {
  apiKey: process.env.REACT_APP_FIREBASE_API_KEY,
  authDomain: process.env.REACT_APP_FIREBASE_AUTH_DOMAIN,
  databaseURL: process.env.REACT_APP_FIREBASE_DATABASE_URL,
  projectId: process.env.REACT_APP_FIREBASE_PROJECT_ID,
  storageBucket: process.env.REACT_APP_FIREBASE_STORAGE_BUCKET,
  messagingSenderId: process.env.REACT_APP_FIREBASE_MESSAGING_SENDER_ID,
  appId: process.env.REACT_APP_FIREBASE_APP_ID,
};

if (!firebase.apps.length) {
  firebase.initializeApp(firebaseConfig)
}

const auth = firebase.auth();
auth.useEmulator("http://localhost:9099");

function App() {
  const [idToken, setIdToken] = useState<string>('')
  const [userId, setUserId] = useState<string>('')

  const queryString = "query { users { id name email } }"
  const query = { query: queryString }

  const login = () => {
    firebase.auth().signInWithPopup(provider)
  }

  const logout = () => {
    firebase.auth().signOut()
    setIdToken("")
    setUserId("")
  }

  const fetchMetadata = () => {
    const metadataRef = firebase.database().ref(`metadata/${userId}/refreshTime`);
    metadataRef.on("value", snapshot => {
      const data = snapshot.val();
      console.dir(data)
    });
  }

  // XXX: Too hacky
  // See: https://github.com/hasura/graphql-engine/issues/6338#issuecomment-762556202
  const getLocallySignedToken = (token: string) => {
    const decoded = jwt.decode(token);
    if (decoded) {
      return jwt.sign(decoded, "secret_secret_secret_secret_secret")
    }
  }

  const isLocalEnvironment = () => {
    return window.location.hostname === "localhost"
  }

  useEffect(() => {
    return firebase.auth().onAuthStateChanged(async user => {
      if (user) {
        setUserId(user.uid)
        const token = await user.getIdToken()
        const idTokenResult = await user.getIdTokenResult()
        const hasuraClaim = idTokenResult.claims['https://hasura.io/jwt/claims']

        if (hasuraClaim) {
          setIdToken(token)
        } else {
          console.log('refresh')
          // Check if refresh is required.
          const metadataRef = firebase.database().ref('metadata/' + user.uid + '/refreshTime')
          metadataRef.on('value', async data => {
            console.log("fetch from realtime-database")
            if (!data.exists) {
              return
            }
            const token = await user.getIdToken(true)
            setIdToken(token)
          })
        }
      } else {
        // out
      }
    })
  }, [])

  const fetchUsers = () => {
    console.log(idToken)
    fetch(process.env.REACT_APP_HASURA_GRAPHQL_ENDPOINT, {
      method: 'POST',
      body: JSON.stringify(query),
      headers: {
        Authorization: `Bearer ${
          isLocalEnvironment() ? getLocallySignedToken(idToken) : idToken
        }`
      }
    }).then(response => {
      response.json().then(result => {
        console.dir(result)
      })
    })
  }

  return (
    <div>
      <div>
        <button onClick={login}>
          Log in
        </button>
      </div>
      <div>
        <button onClick={logout}>
          Log out
        </button>
      </div>
      <div>
        <button onClick={fetchMetadata}>
          Fetch metadata
        </button>
      </div>
      <div>
        <button onClick={fetchUsers} disabled={!idToken.length}>
          fetch
        </button>
        { !!idToken.length && <div>
            {getLocallySignedToken(idToken)}
          </div>
        }
        { !!userId.length && <div>
            userId: {userId}
          </div>
        }
      </div>
      <a href="https://jwt.io/" target="_blank" rel="noreferrer">jwt.io</a>
    </div>
  );
}

export default App;
