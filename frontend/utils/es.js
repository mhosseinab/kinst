var bodybuilder = require('bodybuilder');

export const addElasticFilters = (baseQ, props) => {
  if (props.province) {
    baseQ.query('term', 'province.keyword', props.province);
  }
  if (props.status) {
    baseQ.query('term', 'status.keyword', props.status);
  }
  if (props.city) {
    baseQ.query('term', 'city.keyword', props.city);
  }
  if (props.expert) {
    baseQ.query('term', 'expert_status.keyword', props.expert);
  }
  if (props.damage_type) {
    baseQ.query('term', 'damage_type', props.damage_type);
  }
  if (props.company_id) {
    const compaines = props.company_id.split(',');
    if (compaines.length <= 10) {
      baseQ.query('term', 'company_id.keyword', props.company_id);
    }
  }
  // if (props.from && props.to) {
  //   baseQ.query('term', 'from', props.from);
  //   baseQ.query('term', 'to', props.to);
  // }
};

export const bodyBuilder = () => {
  return bodybuilder();
};
