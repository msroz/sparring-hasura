import React, { useState } from 'react';
import firebase from './firebaseConfig'
import jwt from 'jsonwebtoken'

function App() {
  const [idToken, setIdToken] = useState<string>('')
  const queryString = "query { users { id name email } }"
  const query = { query: queryString }

  const login = () => {
    const provider = new firebase.auth.GoogleAuthProvider()
    firebase.auth().signInWithPopup(provider)
  }

  const logout = () => {
    firebase.auth().signOut()
  }

  // XXX: Too hacky
  // See: https://github.com/hasura/graphql-engine/issues/6338#issuecomment-762556202
  const getLocallySignedToken = token => {
    return jwt.sign(jwt.decode(token), "secret_secret_secret_secret_secret")
  }

  const isLocalEnvironment = () => {
    return window.location.hostname === "localhost"
  }

  firebase.auth().onAuthStateChanged(user => {
    if (user) {
      user.getIdToken().then(token => {
        setIdToken(token)
        console.log(token)
      })
    }
  })

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
        <button onClick={fetchUsers} disabled={!idToken.length}>
          fetch
        </button>
        { !!idToken.length && <div>
            {getLocallySignedToken(idToken)}
          </div>
        }
      </div>
      <a href="https://jwt.io/" target="_blank" rel="noreferrer">jwt.io</a>
    </div>
  );
}

export default App;
