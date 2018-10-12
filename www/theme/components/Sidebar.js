import React from 'react'
import { Link } from 'landr'

export default ({ links }) => {
  {
    links.map(l => {
      return (
        <Box>
          <Link style={{ width: '100%' }} color={'white'} p={2} to={l.to}>
            {l.title}
          </Link>
        </Box>
      )
    })
  }
}
