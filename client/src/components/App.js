import React, {Component} from 'react';
import {Link, Route, Switch} from 'react-router-dom';
import {AuthConsumer} from 'lib/auth/authContext.js';

import './App.css';

import EditPage from 'pages/EditPage.js';

class App extends Component {
  render() {
    return (
      <div className="App">
        <div className="App_header">
          <div className="App_container">
            <AuthConsumer>
              {({info, googleAuth}) => (
                <React.Fragment>
                  <span className="App_brand">
                    <Link to="/">Terse</Link>
                  </span>
                  <button className="App_logout" onClick={() => googleAuth.signOut()}>
                    {info && info.email} | Sign Out
                  </button>
                </React.Fragment>
              )}
            </AuthConsumer>
          </div>
        </div>

        <div className="App_container">
          <Switch>
            <Route exact path="/" component={EditPage} />
            <Route exact path="/:slug" component={EditPage} />
            <Route path="*" component={() => <h2>404: Not Found</h2>} />
          </Switch>
        </div>
      </div>
    );
  }
}

export default App;
