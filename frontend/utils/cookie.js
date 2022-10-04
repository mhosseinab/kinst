import Cookie from 'universal-cookie';

const cookies = new Cookie();

export const setCookie = (name, value) => {
  cookies.set(name, value);
};

export const getCookie = name => {
  return cookies.get(name);
};

export const clearCookie = name => cookies.remove(name);
