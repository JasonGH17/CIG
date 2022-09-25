import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './index.css';
import App from './pages/App';
import Menu from './components/menu/Menu';
import Project from './pages/Project';

const root = ReactDOM.createRoot(
	document.getElementById('root') as HTMLElement
);
root.render(
	<React.StrictMode>
		<BrowserRouter>
			<Routes>
				<Route path="/" element={<Menu />}>
					<Route index element={<App />} />
					<Route path="p" element={<Project />} />
				</Route>
			</Routes>
		</BrowserRouter>
	</React.StrictMode>
);
