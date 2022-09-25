import { useEffect, useState } from 'react';
import getPRJ from '../data/prj';
import useQuery from '../hooks/useQuery';
import type { PRJ } from '../types';

function Project() {
	const query = useQuery();

	const [prj, setprj] = useState<PRJ>({
		name: '',
		refresh: 0,
		location: '',
	});

	useEffect(() => {
		getPRJ().then((data: PRJ[]) => {
			setprj(
				data.find((val) => val.name === query.get('name')) || {
					name: '',
					refresh: 0,
					location: '',
				}
			);
		});
	}, [query]);

	return (
		<div>
			<h2>{prj.name}</h2>
		</div>
	);
}

export default Project;
