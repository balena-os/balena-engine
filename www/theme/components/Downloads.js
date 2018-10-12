import React from 'react';
import _ from 'lodash';
import styled, { withTheme } from 'styled-components';
import { Flex, Box, Container, Text, Link, Heading } from 'resin-components';
import Button from '../components/Button';

import redFill from '../images/red-fill.svg';

function createMarkup(html) {
  return { __html: html };
}

const DownloadsWrapper = styled(Box)`
  background-image: url(${redFill});
  background-repeat: no-repeat;
  background-position: left;
  background-size: contain;
`;

const StyledLink = styled(Link)`
  && {
    color: #fff;
    &:hover {
      color: #fff;
    }
  }
`;

const TableWrapper = styled(Flex)`
  box-shadow: -10px 9px 21px 0 rgba(152, 173, 227, 0.08);
  border: solid 1px ${props => props.theme.colors.primary.main};
  border-radius: ${props => props.theme.radius}px;
  background-color: #ffffff;
  position: relative;

  &:before {
    content: ' ';
    width: 0;
    height: 0;
    border-left: 12px solid transparent;
    border-right: 12px solid transparent;
    border-bottom: 12px solid ${props => props.theme.colors.primary.main};

    position: absolute;
    left: calc(50% - 6px);
    top: -12px;
    z-index: 2;
    }
    &:after {
      content: ' ';
      width: 0;
      height: 0;
      border-left: 12px solid transparent;
      border-right: 12px solid transparent;
      border-bottom: 12px solid #fff;

      position: absolute;
      left: calc(50% - 6px);
      top: -11px;
      z-index: 3;
      }
  }

`;

const Table = styled.table`
  width: 100%;

  > tr {
    > th,
    td {
      text-align: left;
    }
    > th {
      padding: 14px 20px 14px 0;
    }
    > td {
      padding: 20px 20px 20px 0;
      border-top: 1px solid rgba(214, 221, 242, 0.5);
    }
  }
`;

export default withTheme(props => {
  const release = _.get(props, 'releases[0]');
  return (
    <DownloadsWrapper py={120}>
      <Container align="center">
        <Heading.h5
          id="download"
          fontSize={14}
          mb={16}
          color={props.theme.colors.primary.main}
        >
          DOWNLOAD
        </Heading.h5>
        <Heading.h1 fontSize={34} mb={40}>
          Get your assets
        </Heading.h1>
        <TableWrapper justify="center">
          <Box
            width={[1, 1, 1, 2 / 3]}
            py={[40, 40, 40, 80]}
            px={[20, 20, 20, 0]}
          >
            <Table>
              <tr>
                <th>
                  <Heading.h6
                    fontSize={12}
                    color={props.theme.colors.info.main}
                  >
                    ASSET
                  </Heading.h6>
                </th>
                <th>
                  <Heading.h6
                    fontSize={12}
                    color={props.theme.colors.info.main}
                  >
                    OS
                  </Heading.h6>
                </th>
                <th>
                  <Heading.h6
                    fontSize={12}
                    color={props.theme.colors.info.main}
                  >
                    ARCH
                  </Heading.h6>
                </th>
                <th>&nbsp;</th>
              </tr>
              {release &&
                release.assets.map((asset, index) => (
                  <tr key={index}>
                    <td>
                      <Heading.h6 fontSize={14}>{asset.name}</Heading.h6>
                    </td>
                    <td>
                      <Heading.h6 fontSize={14}>{asset.os}</Heading.h6>
                    </td>
                    <td>
                      <Heading.h6 fontSize={14}>{asset.arch}</Heading.h6>
                    </td>
                    <td>
                      <Flex justify="flex-end">
                        <Button round primary>
                          <StyledLink
                            mx={3}
                            blank
                            href={asset.browser_download_url}
                          >
                            Download
                          </StyledLink>
                        </Button>
                      </Flex>
                    </td>
                  </tr>
                ))}
            </Table>
          </Box>
        </TableWrapper>
      </Container>
    </DownloadsWrapper>
  );
});
