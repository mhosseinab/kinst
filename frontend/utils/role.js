import { getCookie } from './cookie';
import { parseJwt } from './jwt';

export const userRoleAdmin = 'ADMIN';
export const userRoleBranch = 'BRANCH';
export const userRoleTavanir = 'TAVANIR';

const roles = {
  userRoleAdmin: 'مدیر سیستم',
  userRoleBranch: 'شعب بیمه',
  userRoleTavanir: 'توانیر',
};

const rolesTitles = {
  ADMIN: 'مدیر سیستم',
  BRANCH: 'شعب بیمه',
  TAVANIR: 'توانیر',
  REPORTER: 'گزارش گیرنده',
};

export const isUserAdmin = () => {
  const cookie = getCookie(authToken);
  if (cookie) {
    const user = parseJwt(cookie);
    if (user.role === userRoleAdmin) {
      return true;
    }
  }
  return false;
};

export const isUserBranch = () => {
  const cookie = getCookie(authToken);
  if (cookie) {
    const user = parseJwt(cookie);
    if (user.role === userRoleBranch) {
      return true;
    }
  }
  return false;
};

const authToken = 'token';
export const getBranchState = () => {
  const cookie = getCookie(authToken);
  if (cookie) {
    const user = parseJwt(cookie);
    return user.state;
  }
  return false;
};

export const isUserTavanir = () => {
  const cookie = getCookie(authToken);
  if (cookie) {
    const user = parseJwt(cookie);
    if (user.role === userRoleTavanir) {
      return true;
    }
  }
  return false;
};

export const getRoleTitle = r => {
  return rolesTitles[r];
};
