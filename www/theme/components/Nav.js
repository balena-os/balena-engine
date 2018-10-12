import React from 'react';
import { withTheme } from 'styled-components';
import {
  Image,
  Box,
  Flex,
  Container,
  Text,
  Link as RLink,
  DropDownButton,
} from 'resin-components';
import styled from 'styled-components';
import balenaLogo from '../images/balena.svg';
import EngineLogo from '../images/balena-engine.svg';
import { Link as RouterLink } from '@resin.io/react-static';

import { navigationLinks } from '../helpers';

import githubIcon from '../images/github.svg';

const windowGlobal = typeof window !== 'undefined' && window;

const GithubLink = styled(RLink)`
  position: absolute;
  top: -16px;
  right: 0;
`;

const Brand = () => (
  <Box>
    <Image style={{ height: '40px' }} src={EngineLogo} />
  </Box>
);

const MenuLink = styled(RLink)`
	&& {
		position: relative;
		font-size: 14px;
		transition: color .1s ease-in;
		font-weight: bold;

		&:before {
			content: ${props => (props.underline ? ' ' : 'none')};
			position: absolute;
			left: 0;
			right: 0;
			bottom: -4px;
			opacity: 0;
			transition: opacity .1s ease-in;
			border: 1px solid ${props => props.theme.colors.primary.main};
		}
		&:hover {
			color: ${props => props.theme.colors.primary.main};
			border-color: ${props => props.theme.colors.primary.main}
			&:before {
				opacity: 1;
			}
		}
	}
`;

const TranparentHeader = styled(Box)`
  && {
    position: absolute;
    top: 0;
    width: 100%;
    background: transparent;
    z-index: 2;
  }
`;

const SolidHeader = styled(Box)`
  position: sticky;
  top: 0px;
  z-index: 2;
  box-shadow: rgb(239, 239, 239) 0px 0px 2px 1px;
  background: rgb(255, 255, 255);
`;

const MobileNavigation = styled(DropDownButton)`
  display: none;
  @media all and (max-width: ${props => props.theme.breakpoints[2]}em) {
    display: block;
  }
  && > div {
    left: -85px;
    box-shadow: none;
  }
`;

const DesktopNavigation = styled(Flex)`
  display: none;
  @media all and (min-width: ${props => props.theme.breakpoints[2]}em) {
    display: block;
  }
`;

const Nav = withTheme(props => {
  let pathSlug = null;
  if (windowGlobal) {
    pathSlug = windowGlobal.location.pathname;
    pathSlug = _.trim(pathSlug, '/');
  }

  const isIndex = _.isEmpty(pathSlug);

  // If Index, show the transparent Header, that blends with the svg background
  const HeaderElement = isIndex ? TranparentHeader : SolidHeader;
  return (
    <HeaderElement py={3}>
      <Container>
        <Flex
          justify="space-between"
          align="center"
          pt={2}
          style={{ position: 'relative' }}
        >
          <Box>
            <Flex wrap mb={1} align="center">
              <Text.span fontSize={12}>Open source project by</Text.span>
              <RLink blank href="https://balena.io">
                <Image ml={2} style={{ height: '20px' }} src={balenaLogo} />
              </RLink>
            </Flex>
            <RouterLink mt={2} to="/">
              <Brand />
            </RouterLink>
          </Box>

          <MobileNavigation joined secondary outline>
            <Flex direction="column">
              {navigationLinks.map((entry, i) => (
                <MenuLink
                  color={props.theme.colors.text.main}
                  my={2}
                  key={i}
                  href={entry.link}
                  to={entry.link}
                  is={!entry.withHash ? RouterLink : 'a'}
                >
                  {entry.text}
                </MenuLink>
              ))}
            </Flex>
          </MobileNavigation>
          <DesktopNavigation align="center">
            {navigationLinks.map((entry, i) => (
              <MenuLink
                color={props.theme.colors.text.main}
                underline
                ml={4}
                key={i}
                href={entry.link}
                to={entry.link}
                is={!entry.withHash ? RouterLink : 'a'}
              >
                {entry.text}
              </MenuLink>
            ))}
          </DesktopNavigation>

          <GithubLink href={props.repository.html_url} blank>
            <Image src={githubIcon} />
          </GithubLink>
        </Flex>
      </Container>
    </HeaderElement>
  );
});

export default Nav;
