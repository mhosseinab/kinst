const { create } = require('apisauce');
const Cookies = require('universal-cookie');

const baseURL = process.env.HOSTURL || 'https://example.com';
// create main request configs
const request = create({
  baseURL,
  headers: {
    'Content-Type': 'application/json',
    'Accept-Language': 'fa',
  },
});

// server middleware used to add required headers to request
const serverRequestModifier = req => {
//   const cookies = new Cookies(req.headers.cookie);
//   const accessToken = cookies.get('token');
//   const authorization = accessToken ? accessToken : null; 
//   if (authorization) {
//     request.setHeader('Authorization', authorization);
//   }
//   return request
};

// create login configs
const loginConfig = {
  'Content-Type': 'application/x-www-form-urlencoded',
};

module.exports = {
  request,
  baseURL,
  loginConfig,
  serverRequestModifier,
};
