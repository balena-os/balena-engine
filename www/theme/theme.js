import { lighten, darken } from 'resin-components/dist/utils';
import reduce from 'lodash/reduce';
import get from 'lodash/get';
import isObject from 'lodash/isObject';

const primary = '#eb407a';
const secondary = '#f0e9f3';
const tertiary = '#c1c7dd';
const quartenary = '#f8f9fd';

const danger = '#FF423D';
const warning = '#FCA321';
const success = '#76C950';
const info = '#527699';
const text = '#2a506f';
const gray = {
  light: '#f4f4f4',
  main: '#c6c8c9',
  dark: '#292b2c',
};

const createShades = color => {
  if (isObject(color)) return color;
  return {
    main: color,
    light: lighten(color),
    dark: darken(color),
  };
};

const defaultColors = {
  primary,
  secondary,
  tertiary,
  quartenary,
  danger,
  warning,
  success,
  info,
  text,
  gray,
};

const createColors = colors => {
  return reduce(
    colors,
    (acc, val, key) => {
      acc[key] = createShades(val);
      return acc;
    },
    {},
  );
};

const defaultControlHeight = 36;
const emphasizedControlHeight = 45;

export const breakpoints = [32, 48, 64, 90];

export const space = [
  0,
  4,
  8,
  16,
  defaultControlHeight,
  emphasizedControlHeight,
  128,
];

export const fontSizes = [12, 14, 16, 20, 24, 32, 48, 64, 72, 96];

export const weights = [400, 700];

export const radius = 6;
export const font = 'Nunito, Arial, sans-serif';
export const monospace = "'Ubuntu Mono', 'Courier New', monospace";

const theme = userTheme => {
  return {
    colors: createColors({
      ...defaultColors,
      ...get(userTheme, 'colors'),
    }),
    breakpoints,
    space,
    fontSizes,
    weights,
    font,
    monospace,
    radius,
  };
};

const globalStyles = theme => `

@font-face {
  font-family: "Nunito";
	src: url("/fonts/Nunito-Regular.ttf") format("ttf");
	font-weight: normal;
}

@font-face {
  font-family: "Nunito";
	src: url("/fonts/Nunito-Regular.ttf") format("ttf");
	font-weight: bold;
}

@font-face {
  font-family: "CircularStd";
	src: url("/fonts/CircularStd-Book.otf") format("opentype");
	font-weight: bold;
}


* { box-sizing: border-box; }
body {
	font-family: Nunito;
	margin: 0;
	color: ${theme.colors.text.main};
}
p {
	line-height: 1.5;
	> em {
		border-left: 4px solid ${theme.colors.primary.main}90;
		display: block;
		padding-left: ${theme.space[2]}px;
		background: ${theme.colors.primary.main}20;
		border-radius: ${theme.radius}px;
		padding: ${theme.space[3]}px;
		color: #333;
	}
}
ul {
  margin-bottom: ${theme.space[4]}px;
  li {
    margin-bottom: ${theme.space[2]}px;
  }
}
img { max-width: 100%; }
a {
  color: ${theme.colors.primary.main};
  text-decoration: none;
  cursor: pointer;
}
blockquote {
  font-style: italic;
  padding: 20px;
  > p {
    box-shadow: inset 0 -5px 0px 0px ${theme.colors.primary.main};
  }
}
code {
	color: ${theme.colors.text.main};
	border-radius: ${theme.radius}px;
}
`;

export default theme;
export { globalStyles };
