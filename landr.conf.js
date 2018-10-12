const getArch = str => {
  const [_, arch] = str.match(/-([^-]+)\.tar\.gz$/);
  return arch;
};

const packagePrettyName = str => `Balena for ${getArch(str)}`;

const prepAssets = release => {
  release.assets = release.assets.map(asset => {
    return Object.assign({}, asset, {
      prettyName: packagePrettyName(asset.name),
      arch: getArch(asset.name),
      os: 'Linux',
    });
  });

  return release;
};

module.exports = {
  plugins: [
    {
      middleware: (store, action, next) => {
        if (action.type === 'ADD_RELEASE') {
          // intercept all releases and add pretty labels to assets
          action.payload = prepAssets(action.payload);
        }

        return next(action);
      },
    },
  ],
  theme: 'www/theme',
  settings: {
    title: 'balenaEngine',
    lead: 'A Moby-based container engine for IoT',
    description:
      'BalenaEngine is a new container engine purpose-built for embedded and IoT use cases and compatible with Docker containers.',
    analytics: {
      mixpanelToken: '81dd4ca517f8bd50d46aa8fe057882ac', // mixpanelToken
      gosquaredId: 'GSN-122115-A', // gosquared Id
      gaSite: 'balena.io', // google Analytics site eg balena.io
      gaId: 'UA-45671959-5', // google Analytics ID
    },
    installCommand: 'curl -sfL https://balena.io/install.sh | sh',
    featuresLead: 'An engine compatible with docker containers',
    features: [
      {
        title: 'Small footprint',
        icon: 'footprint',
        description: '3.5x smaller than Docker CE, packaged as a single binary',
      },
      {
        title: 'Multi-arch support',
        icon: 'multiple',
        description:
          'Available for a wide variety of chipset architectures, supporting everything from tiny IoT devices to large industrial gateways',
      },
      {
        title: 'True container deltas',
        icon: 'bandwidth',
        description:
          'Bandwidth-efficient updates with binary diffs, 10-70x smaller than pulling layers in <a href="https://balena.io/blog/announcing-balena-a-moby-based-container-engine-for-iot/#technical_comparison" target="_blank">common scenarios</a>',
      },
      {
        title: 'Minimal wear-and-tear',
        icon: 'storage',
        description:
          'Extract layers as they arrive to prevent excessive writing to disk, protecting your storage from eventual corruption',
      },
      {
        title: 'Failure-resistant pulls',
        icon: 'failure-resistant',
        description:
          'Atomic and durable image pulls defend against partial container pulls in the event of power failure',
      },
      {
        title: 'Conservative memory use',
        icon: 'undisturbed',
        description:
          'Prevents page cache thrashing during image pull, so your application runs undisturbed in low-memory situations',
      },
    ],
    motivation: {
      blogPost:
        'https://balena.io/blog/announcing-balena-a-moby-based-container-engine-for-iot',
      paragraphs: [
        'balenaEngine is a new container engine purpose-built for embedded and IoT use cases and compatible with Docker containers. Based on Docker’s Moby Project, balenaEngine supports container deltas for 10-70x more efficient bandwidth usage, has 3.5x smaller binaries, uses RAM and storage more conservatively, and focuses on atomicity and durability of container pulling.',
        'Having seen IoT devices used in production for tens of millions of hours, we’ve become intimately acquainted with the unique needs of the embedded world. Until recently, we addressed these by either making small modifications to Docker itself or building larger components outside of it, but that approach had reached its limits. Meanwhile, as the Docker binaries have grown in functionality, they have also grown in size, eating away.',
      ],
    },
  },
};
