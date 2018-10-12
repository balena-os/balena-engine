import React from 'react';
import { Flex, Box, Text, Image, Link as RLink } from 'resin-components';
import { Link as RouterLink } from 'landr';
import styled, { withTheme } from 'styled-components';

import balenaLogo from '../images/balena.svg';
import engineLogo from '../images/balena-engine.svg';

import { navigationLinks } from '../helpers';

const MenuLink = styled(RLink)`
	&& {
		color: ${props => props.theme.colors.text.main};
		position: relative;
		font-size: 14px;
		font-weight: bold;
		transition: color .1s ease-in;

		&:hover {
			color: ${props => props.theme.colors.primary.main};
			border-color: ${props => props.theme.colors.primary.main}
			&:before {
				opacity: 1;
			}
		}
	}
`;

const Footer = ({ repository, ...props }) => {
  return (
    <Box py={60} mt={60} bg={props.theme.colors.quartenary.main}>
      <Flex align="center" direction="column" justify="center" wrap>
        <Box mb={50}>
          {navigationLinks.map((entry, i) => (
            <MenuLink
              underline
              mx={3}
              key={i}
              blank={entry.isBlank}
              href={entry.link}
              to={entry.link}
              is={!entry.withHash ? RouterLink : 'a'}
            >
              {entry.text}
            </MenuLink>
          ))}
        </Box>
        <Flex align="center" justify="center" w="100%">
          <Image mr={3} style={{ height: '30px' }} src={engineLogo} />
          <Text fontSize={12}>An open source project by</Text>
          <RLink blank href="https://resin.io">
            <Image ml={3} style={{ height: '20px' }} src={balenaLogo} />
          </RLink>
        </Flex>
      </Flex>
    </Box>
  );
};

export default withTheme(Footer);
