import React from 'react';
import { Box, Container, Flex, Text, Divider, Heading } from 'resin-components';
import Accordian from 'components/Accordian';
import { withTheme } from 'styled-components';

const FAQ = props => {
	const items = props.faqs.map(faq => {
		return {
			title: (
				<Text my={20} fontSize={16} bold>
					{faq.title}
				</Text>
			),
			render: () => {
				return (
					<Text.p
						fontSize={14}
						mt={10}
						mb={20}
						dangerouslySetInnerHTML={{ __html: faq.html }}
					/>
				);
			},
		};
	});

	return (
		<Container py={5}>
			<Flex justify="center">
				<Box width={5 / 6}>
					<Heading.h5 fontSize={14} mb={16} color={props.theme.colors.primary.main}>
						BALENA FIN
					</Heading.h5>
					<Heading.h1 fontSize={34} mb={60}>
						Frequently asked questions
					</Heading.h1>
					<Divider m={0} height={1} color={props.theme.colors.tertiary.main} />
					<Accordian items={items} />
				</Box>
			</Flex>
		</Container>
	);
};

export default withTheme(FAQ);
