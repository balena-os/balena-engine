import React, { Component } from 'react'
import Sniffr from 'sniffr'
import sortBy from 'lodash/sortBy'
import arch from 'arch'
import { DropDownButton, Text } from 'resin-components'
import { withTheme } from 'styled-components'
import { Link } from 'landr'
import get from 'lodash/get'

const Asset = ({ asset, color, ...props }) => {
  // TODO make PR width prop into resin-components lib
  return (
    <Link
      style={{ width: '100%' }}
      {...props}
      color={color}
      to={asset.browser_download_url}
    >
      {asset.prettyName || asset.name}
    </Link>
  )
}

class DownloadButton extends Component {
  constructor(props) {
    super(props)
    this.state = {
      primaryAsset: null,
      assets: props.assets
    }
  }

  componentDidMount() {
    const assets = this.state.assets
    // only run if we have assets and if there is an os prop
    if (assets.length < 1 || !assets[0].os) {
      this.setState({
        primaryAsset: null,
        assets: assets
      })
      return
    }

    const client = new Sniffr()
    client.sniff(window.navigator.userAgent)
    client.os.arch = arch()

    // give points for not matching
    const score = (condition, p) => (!condition ? p : 0)

    const sortedAssets = sortBy(assets, l => {
      let assetScore = score(
        l.os.toLowerCase() === client.os.name.toLowerCase(),
        2
      )
      if (assetScore === 0) {
        assetScore = assetScore + (l.arch === client.os.arch, 1)
      }

      return assetScore
    })

    this.setState({
      primaryAsset: sortedAssets.shift(),
      assets: sortedAssets
    })
  }

  render(props) {
    const getter = key => get(props, key)
    const tracker = this.context.tracker
    const assets = [...this.state.assets].filter(t => {
      // etcher specifc code
      return t.type !== 'CLI'
    })

    return (
      <DropDownButton
        {...props}
        emphasized
        primary
        joined={!this.state.primaryAsset}
        label={
          this.state.primaryAsset ? (
            <Asset
              px={3}
              onClick={() => {
                tracker.create('download', this.state.primaryAsset)
              }}
              asset={this.state.primaryAsset}
              color={'white'}
            />
          ) : (
            <Text px={3}>Download</Text>
          )
        }
      >
        {assets.map((asset, i) => {
          return (
            <Asset
              key={i}
              py={2}
              asset={asset}
              onClick={() => {
                tracker.create('download', asset)
              }}
              color={getter('theme.colors.gray.dark')}
            />
          )
        })}
      </DropDownButton>
    )
  }
}

DownloadButton.contextTypes = {
  tracker: React.PropTypes.object
}

export default withTheme(DownloadButton)
