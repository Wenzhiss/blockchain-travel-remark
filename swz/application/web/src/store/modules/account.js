import { login } from '@/api/account';
import { getToken, setToken, removeToken } from '@/utils/auth';
import { resetRouter } from '@/router';

const getDefaultState = () => {
	return {
		token: getToken(),
    storeID:'',
		name: '',
	};
};

const state = getDefaultState();

const mutations = {
	RESET_STATE: (state) => {
		Object.assign(state, getDefaultState());
	},
	SET_TOKEN: (state, token) => {
		state.token = token;
	},
	SET_STOREID: (state, storeID) => {
		state.storeID = storeID;
	},
	SET_NAME: (state, name) => {
		state.name = name;
	},
};

const actions = {
	login({ commit }, storeID) {
		return new Promise((resolve, reject) => {
			login({
				args: [
					{
						storeID: storeID,
					},
				],
			})
				.then((response) => {
					commit('SET_TOKEN', response[0].storeID);
					setToken(response[0].storeID);
					resolve();
				})
				.catch((error) => {
					reject(error);
				});
		});
	},
	// get user info
	getInfo({ commit, state }) {
		return new Promise((resolve, reject) => {
			login({
				args: [
					{
						storeID: state.token,
					},
				],
			})
				.then((response) => {
					commit('SET_STOREID', response[0].storeID);
					commit('SET_NAME', response[0].name);
				})
				.catch((error) => {
					reject(error);
				});
		});
	},
	logout({ commit }) {
		return new Promise((resolve) => {
			removeToken();
			resetRouter();
			commit('RESET_STATE');
			resolve();
		});
	},

	resetToken({ commit }) {
		return new Promise((resolve) => {
			removeToken();
			commit('RESET_STATE');
			resolve();
		});
	},
};

export default {
	namespaced: true,
	state,
	mutations,
	actions,
};
