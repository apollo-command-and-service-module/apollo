import React, { Component } from "react";
import "./SyncHistory.scss";

class SyncHistory extends Component {
    render() {
        const messages = this.props.syncHistory.map((msg, index) => (
            <p key={index}>{msg.data}</p>
        ));

        return (
            <div className="SyncHistory">
                <h2>Sync History</h2>
                {messages}
            </div>
        );
    }
}

export default SyncHistory;