import React from 'react'

// TODO use state to display a collection of messages
function MessageComponent(props){
	let messages = props.messages.map((item, index) => {return (<Message>{item}</Message>);});
	return(
		<div className="messageComponent">
			{ messages }
		</div>
	);
}

function Message(props){
	return(
		<div className="message">
			{ props.children }
		</div>
	);
}

export default MessageComponent;