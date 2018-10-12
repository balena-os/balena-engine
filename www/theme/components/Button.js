import { Button as RButton } from 'resin-components';
import styled from 'styled-components';

const Button = styled(RButton)`
	font-size: 14px;
	font-weight: bold;
	border-radius: ${props => (props.round ? '20px' : 'inherit')};
	border: ${props => (props.borderless ? 'none' : 'auto')};
	background: ${props => (props.outline && props.primary ? '#fff' : 'auto')};

	&:hover {
		> a {
			color: ${props => (props.outline ? '#fff' : 'auto')};
		}
	}
`;

export default Button;
