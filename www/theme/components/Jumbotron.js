import React from 'react';
import styled, { withTheme } from 'styled-components';
import get from 'lodash/get';
import {
  Container,
  Heading,
  Banner,
  Image,
  Box,
  Flex,
  Text,
  Link as RLink,
} from 'resin-components';
import { Link } from 'landr';
import Button from '../components/Button';

import bg from '../images/hero-bg.svg';

const HeroBanner = styled(Banner)`
  display: block;
  background-image: url(${bg});
  background-position: bottom;
  background-repeat: no-repeat;
  background-size: 100% -10%;

  @media all and (max-width: ${props => props.theme.breakpoints[2]}em) {
    background-image: linear-gradient(
      157deg,
      rgba(255, 255, 255, 0.5),
      #ff80ab 24%,
      #eb407a
    );
    text-align: center;
    padding-bottom: 24px;
  }
`;

const Code = styled.code`
@keyframes blinker {
  50% {
    opacity: 0;
  }
}

  text-align: left;
  line-height: 1.4;
  display: inline-block;
  font-weight: bold;
  background-color: #ffffff30;
  font-size: 14px;
  padding: 16px 20px;
  position: relative;

  &:after {
    content: ' ';
    animation: blinker 1s linear infinite;
    border-left: 2px solid ${props => props.theme.colors.primary.main} ;
  }
`;

const linkToDownload = () => {
  window.location.hash = '';
  window.location.hash = '#download';
};

export default withTheme(props => {
  const getter = key => get(props, key);
  const version = getter('releases[0]');
  return (
    <HeroBanner>
      <Container pt={140} pb={80}>
        <Flex justify="center">
          <Box width={[1, 1, 1, 5 / 6]}>
            <Box width={[1, 1, 1, 7 / 10]} mb={[30, 30, 0]}>
              <Heading.h1 fontSize={[30, 35, 45, 52]} mb={25}>
                {getter('settings.lead')}
              </Heading.h1>
            </Box>
            <Flex mx={-10} wrap>
              <Box width={[1, 1, 1, 1 / 2]} px={10} mb={40}>
                <Text fontSize={16} mb={25} style={{ lineHeight: '1.6' }}>
                  {getter('settings.description')}
                </Text>
                <Box mb={25}>
                  <Code>{getter('settings.installCommand')}</Code>
                </Box>

                <Flex justify={['center', 'center', 'center', 'flex-start']} align='center'>
                  <Button
                    round
                    primary
                    outline
                    borderless
                    px={5}
                    mr={16}
                    onClick={linkToDownload}
                  >
                    Or download
                  </Button>
                  <Text fontSize={12}>
                    {version.tag_name}{' '}
                    <Button underline text ml={8}>
                      <Link
                        color={props.theme.colors.text.main}
                        to="/changelog"
                      >
                        See what's new
                      </Link>
                    </Button>
                  </Text>
                </Flex>
              </Box>
              <Box width={[1, 1, 1, 1 / 2]} px={10}>
                Image
              </Box>
            </Flex>
          </Box>
        </Flex>
      </Container>
    </HeroBanner>
  );
});
