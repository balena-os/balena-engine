import React from 'react'
import { Container, Box } from 'resin-components'

function createMarkup(html) {
  return { __html: html }
}

export default ({ html }) => {
  return (
    <Container pt={2} pb={5}>
      <Box dangerouslySetInnerHTML={createMarkup(html)} />
    </Container>
  )
}
