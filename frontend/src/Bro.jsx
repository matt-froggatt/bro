import React from 'react';
import messages from './App';

function BroComponent(props) {
	return(
		<div className="BroComponent">
			<BroButton ws={props.ws} messages={props.messages} setMessages={props.setMessages}></BroButton>
		</div>
	);
}

function BroButton(props) {
	return(
		<button className="BroButton" onClick={() => props.setMessages(props.ws, props.messages, "Bro")}>Bro</button>
	);
}

export default BroComponent;
