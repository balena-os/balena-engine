import React from 'react';
import { Router, Routes, getSiteProps, Tracker } from 'landr';
import { Provider } from 'resin-components';
import styled, { injectGlobal } from 'styled-components';
import get from 'lodash/get';

import Nav from './components/Nav';
import Footer from './components/Footer';
import Helmet from './components/Helmet';
import ThemeStyles, { globalStyles } from './theme';

const Wrapper = styled.div`
	display: flex;
	min-height: 100vh;
	flex-direction: column;
`;

const Content = styled.div`
	flex: 1;
	display: flex;
	flex-direction: column;
`;

export default getSiteProps(props => {
	const getProp = key => get(props, key);
	const mergedTheme = ThemeStyles(getProp('settings.theme'));
	injectGlobal`${globalStyles(mergedTheme)}`;
	return (
		<Router>
			<Tracker
				prefix={getProp('repository.name')}
				analytics={getProp('settings.analytics')}
			>
				<Provider theme={mergedTheme}>
					<Wrapper>
						<Helmet {...props} />
						<Nav {...props} />
						<Content>
							<Routes />
						</Content>
						<Footer {...props} />
					</Wrapper>
				</Provider>
			</Tracker>
		</Router>
	);
});
