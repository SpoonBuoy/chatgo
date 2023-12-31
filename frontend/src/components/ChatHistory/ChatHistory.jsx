import React from 'react';
import './ChatHistory.scss';

const ChatHistory = ({ chatHistory }) => {
  const messages = chatHistory.map((msg, idx) => (
    <p key={idx}> {msg.data} </p>
  ));

  return (
    <div className="ChatHistory">
      <h2>Chat History</h2>
      {messages}
    </div>
  );
};

export default ChatHistory;

