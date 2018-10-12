import React from 'react';
import { Container } from 'resin-components';
import Doc from '../components/Doc';
import { getSiteProps } from '@resin.io/react-static';

export default getSiteProps(props => {
  return (
    <Container>
      <Doc {...props.changelog} />
    </Container>
  );
});
