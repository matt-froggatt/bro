import React from 'react';
import BroComponent from './Bro'
import './App.css';
import MessageComponent from './Messages';

let messages = [[false, "Bro"], [true, "Bro"]];

function App() {
  return (
    <div className="App">
      <MessageComponent messages={messages}></MessageComponent>
      <BroComponent></BroComponent>
    </div>
  );
}

export default App;
