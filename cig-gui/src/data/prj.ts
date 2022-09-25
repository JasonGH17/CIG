const getPRJ = async () => {
	const res = await fetch('/projects', { method: 'GET' });
	return await res.json();
};

export default getPRJ;
