import { API_URL } from '$lib/api.js';

export const load = async ({}) => {
	const res = await fetch(API_URL + 'products', {
		method: 'GET'
	});

	if (res.status !== 200) {
		return { error: true, message: 'Error occured', data: null };
	}

	const data = await res.json();

	return { error: false, message: 'Successfully got data', data: data };
};
