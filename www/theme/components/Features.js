import React from 'react';
import styled, { withTheme } from 'styled-components';

import { Heading, Container, Flex, Box, Image, Text } from 'resin-components';
import { assets } from 'landr';
import whiteFill from '../images/white-fill.svg';

const FeaturesBox = styled(Box)`
	background-image: url(${whiteFill});
	background-repeat: no-repeat;
	background-position: right;
	background-size: contain;
`;

export default withTheme(props => {
	const list = props.settings.features.map((feature, index) => {
		const imageSet = `${assets[`${feature.icon}@1x`]},
		${assets[`${feature.icon}@1x`]} 1x,
		${assets[`${feature.icon}@2x`]} 2x,
		${assets[`${feature.icon}@3x`]} 3x`;
		return (
			<Flex
				key={index}
				px={10}
				order={index + 1}
				width={[1, 1 / 2, 1 / 2, 1 / 4]}
				align="flex-start"
				mb={20}
			>
				<Image
					mr={16}
					h="50px"
					w="45px"
					src={assets[`${feature.icon}`]}
					srcSet={imageSet}
				/>
				<Box mt={8}>
					<Text mb={2} fontSize={18} bold>
						{feature.title}
					</Text>
					<Text.p
						fontSize={16}
						dangerouslySetInnerHTML={{ __html: feature.description }}
					/>
				</Box>
			</Flex>
		);
	});

	return (
		<FeaturesBox py={120}>
			<Container>
				<Flex wrap mb={5} mx={-10}>
					<Flex
						px={10}
						order={[0, 0, 0, 3]}
						width={[1, 1, 1, 1 / 2]}
						mb={[50, 50, 50, 20]}
					>
						<Box width={[1, 1, 1, 5 / 6]} pl={[0, 0, 0, 64]}>
							<Heading.h5
								fontSize={14}
								mb={16}
								color={props.theme.colors.primary.main}
							>
								FEATURES
							</Heading.h5>
							<Heading.h1 fontSize={34}>
								{props.settings.featuresLead}
							</Heading.h1>
						</Box>
					</Flex>
					{list}
				</Flex>
			</Container>
		</FeaturesBox>
	);
});
