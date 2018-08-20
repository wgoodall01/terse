import React from 'react';
import PropTypes from 'prop-types';
import './LinkEdit.css';

import Button from 'components/Button.js';

class LinkEdit extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      short: props.link.short,
      long: props.link.long
    };
  }

  static get propTypes() {
    return {
      link: PropTypes.object.isRequired,
      onSubmit: PropTypes.func.isRequired
    };
  }

  submit = e => {
    const {short, long} = this.state;
    this.props.onSubmit({short, long});
  };

  render() {
    const {short, long} = this.state;
    return (
      <div className="LinkEdit">
        <input
          className="LinkEdit_input LinkEdit_short"
          value={short}
          onChange={e => this.setState({short: e.target.value})}
        />
        <input
          className="LinkEdit_input LinkEdit_long"
          value={long}
          onChange={e => this.setState({short: e.target.value})}
        />
        <div className="LinkEdit_button-row">
          <Button lg onClick={this.submit}>
            Submit
          </Button>
        </div>
      </div>
    );
  }
}

export default LinkEdit;
