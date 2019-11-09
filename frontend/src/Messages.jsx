import React from 'react'

// TODO use state to display a collection of messages
function MessageComponent(props){
	let messages = props.messages.map((item, index) => {return (<Message key={index.toString()}>{item}</Message>);});
	return(
		<div className="MessageComponent">
			{ messages }
		</div>
	);
}

function Message(props){
	let message = (props.children[0]) ? <SentMessage>{ props.children[1] }</SentMessage> : <RecievedMessage>{ props.children[1] }</RecievedMessage>
	return(
		<div className="Message">
			{ message }
		</div>
	);
}

function SentMessage(props){
	return(
		<div className="SentMessage">
			{ props.children }
		</div>
	);
}

function RecievedMessage(props){
	return(
		<div className="RecievedMessage">
			{ props.children }
		</div>
	);
}

export default MessageComponent;