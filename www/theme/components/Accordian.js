import React, { Component } from 'react';
import { Box, Divider, Flex, Text } from 'resin-components';

const Collapse = Box.extend`
	overflow: hidden;
	max-width: 800px;
	max-height: ${props => (props.isOpen ? '100vh' : '0')};
	transition: max-height 0.4s ease-in-out;
`;

const Wrapper = Box.extend`
	cursor: pointer;
`;

class Accordian extends Component {
	constructor() {
		super();
		this.state = {
			openIndex: null,
		};
	}

	toggle(key) {
		if (this.state.openIndex === key) {
			key = null;
		}
		this.setState({
			openIndex: key,
		});
	}

	render() {
		return (
			<Box>
				{this.props.items.map((item, i) => {
					const isOpen = this.state.openIndex === i;
					return (
						<Wrapper
							key={i}
							onClick={() => {
								this.toggle(i);
							}}
						>
							<Flex align="center" justify="space-between">
								<Box width={11 / 12} px={16}>{item.title}</Box>
								<Box width={1 / 12}>
									<Text pr={16} fontSize={24} align="end">
										{isOpen ? '−' : '+'}
									</Text>
								</Box>
							</Flex>
							<Collapse pl={15} isOpen={isOpen}>{item.render()}</Collapse>
							<Divider m={0} height={1} color="#c1c7dd" />
						</Wrapper>
					);
				})}
			</Box>
		);
	}
}

export default Accordian;
