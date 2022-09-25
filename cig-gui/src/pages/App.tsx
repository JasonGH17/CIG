import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import type { PRJ } from '../types';
import Project from '../components/project/Project';
import styles from './App.module.css';
import getPRJ from '../data/prj';

function App() {
	const [prj, setprj] = useState<PRJ[]>([]);

	useEffect(() => {
		getPRJ().then(setprj);
	}, []);

	const nav = useNavigate();

	return (
		<div>
			<h1>Your projects:</h1>

			{(prj.length > 0 && (
				<table className={styles.table}>
					<thead>
						<tr>
							<th style={{ width: '300px' }}>Project Name:</th>
							<th style={{ width: '500px' }}>Last Refresh:</th>
							<th style={{ width: '500px' }}>
								Project Location:
							</th>
							<th style={{ width: '40px' }}></th>
						</tr>
					</thead>
					<tbody>
						{prj.map((val: PRJ) => (
							<Project val={val} nav={nav} />
						))}
					</tbody>
				</table>
			)) || <h2>No Projects are using CIG</h2>}
		</div>
	);
}

export default App;
