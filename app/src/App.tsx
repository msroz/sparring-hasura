import React, { useState } from 'react';
import firebase from './firebaseConfig'

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

  firebase.auth().onAuthStateChanged(user => {
    if (user) {
      user.getIdToken().then(token => {
        setIdToken(token)
        console.log(token)
      })
    }
  })

console.log(process.env.REACT_APP_HASURA_GRAPHQL_ENDPOINT)
  const fetchUsers = () => {
    console.log(idToken)
    fetch(process.env.REACT_APP_HASURA_GRAPHQL_ENDPOINT, {
      method: 'POST',
      body: JSON.stringify(query),
      headers: {
        Authorization: `Bearer ${idToken}`
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
            {idToken}
          </div>
        }
      </div>
      <a href="https://jwt.io/" target="_blank" rel="noreferrer">jwt.io</a>
    </div>
  );
}

export default App;
