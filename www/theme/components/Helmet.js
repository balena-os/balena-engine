import React from 'react'
import { withTheme } from 'styled-components'
import { Head, assets } from 'landr'
import get from 'lodash/get'

export default withTheme(props => {
  const getter = key => get(props, key)
  return (
    <Head>
      <meta charSet="UTF-8" />
      <title>{`${getter('settings.lead') ||
        getter('repository.description')}`}</title>
      <link rel="icon" href={assets['favicon']} type="image/x-icon" />
      <link
        rel="stylesheet"
        href="//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/styles/default.min.css"
      />
      <meta name="theme-color" content={getter('theme.colors.primary.main')} />
      <meta name="og:type" content="og:website" />
      <meta name="og:site_name" content={getter('settings.title')} />
      <meta name="og:description" content={getter('repository.description')} />
      <meta name="og:image" content={assets[getter('repository.name')]} />
    </Head>
  )
})
