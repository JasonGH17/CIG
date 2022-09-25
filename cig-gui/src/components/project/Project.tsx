import type { NavigateFunction } from 'react-router-dom';
import type { PRJ } from '../../types';
import styles from './Project.module.css';
import { AiOutlineBars } from 'react-icons/ai';

function Project({ val, nav }: { val: PRJ; nav: NavigateFunction }) {
	return (
		<tr className={styles.project}>
			<td>
				<h3>{val.name}</h3>
			</td>
			<td>
				<p>
					Last Reload:
					<br />
					{new Date(val.refresh).toLocaleString()}
				</p>
			</td>
			<td>{val.location}</td>
			<td>
				<AiOutlineBars
					className={styles.prop}
					size={25}
					onClick={() =>
						nav({ pathname: '/p', search: `?name=${val.name}` })
					}
				/>
			</td>
		</tr>
	);
}

export default Project;
