import React, { Component } from "react";
import './App.css';
import { connect, sendMsg } from "./api";
import Header from './components/Header/Header';
import SyncHistory from './components/SyncHistory/SyncHistory';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      syncHistory: []
    }
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New Message")
      this.setState(prevState => ({
        syncHistory: [...this.state.syncHistory, msg]
      }))
      console.log(this.state);
    });
  }

  send() {
    console.log("hello");
    sendMsg("hello");
  }

  render() {
    return (
      <div className="App">
        <Header />
        <SyncHistory syncHistory={this.state.syncHistory} />
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}

export default App;
