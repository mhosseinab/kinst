import { shallow } from 'enzyme';
import React from 'react';
import renderer from 'react-test-renderer';

import CPCard from '../CPCard';

describe('CPCard Test', () => {
  it('CPCard title props test"', () => {
    const wrapper = shallow(<CPCard title="Events" />);

    expect(wrapper.prop('title')).toEqual('Events');
  });
});

describe('CPCard Snapshot Testing', () => {
  it('CPCard Snapshot"', () => {
    const component = renderer.create(<CPCard />);
    const tree = component.toJSON();
    expect(tree).toMatchSnapshot();
  });
});
