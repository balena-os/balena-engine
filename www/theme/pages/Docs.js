import React from 'react';
import _ from 'lodash';
import { getSiteProps } from '@resin.io/react-static';
import { Container } from 'resin-components';
import Doc from '../components/Doc';

export default getSiteProps(({ docs }) => {
  return (
    <Container>
      <Doc {...docs[0]} />
    </Container>
  );
});
