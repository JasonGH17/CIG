import { Outlet, useNavigate } from 'react-router-dom';
import styles from './Menu.module.css';

function Menu() {
	const nav = useNavigate();

	return (
		<>
			<div className={styles.topbar}>
				<h1 className={styles.title} onClick={() => nav({ pathname: '/' })}>CIG</h1>
			</div>
			<div className={styles.outlet}>
				<Outlet />
			</div>
		</>
	);
}

export default Menu;
