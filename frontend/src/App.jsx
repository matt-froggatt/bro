import React, {useState} from 'react';
import BroComponent from './Bro'
import './App.css';
import MessageComponent from './Messages';

function App() {
  const [messages, setMessages] = useState([]);

  let ws = new WebSocket("ws://" + window.location.host + "/ws")

  ws.addEventListener("messsage", (e) => {
    var msg = JSON.parse(e.data);
    setMessages(messages.concat[[false, msg.message]]);
  })

  let sendMessage = (ws, messages, message) => {
      ws.send(JSON.stringify({
        Message: message,
      }))
      setMessages(messages.concat([[true, message]]))
  }

  return (
    <div className="App">
      <MessageComponent messages={messages}></MessageComponent>
      <BroComponent ws={ws} messages={messages} setMessages={sendMessage}></BroComponent>
    </div>
  );
}

export default App;
