import React from 'react';
import PropTypes from 'prop-types';
import classnames from 'classnames';
import {Link} from 'react-router-dom';
import './LinkMenu.css';

const LinkMenu = ({links, onSelect, current}) => (
  <React.Fragment>
    {links.map(link => (
      <Link key={link.id} className="LinkMenu_a" to={`/${link.short}`}>
        <div className={classnames('LinkMenu_item', {LinkMenu_selected: link.id === current})}>
          <h2>
            <span className="LinkMenu_shade-span">{window.location.host + '/'}</span>
            <span>{link.short}</span>
          </h2>
          <p>
            <span className="LinkMenu_shade-span">{'to '}</span>
            {link.long}
          </p>
        </div>
      </Link>
    ))}
  </React.Fragment>
);

LinkMenu.propTypes = {
  links: PropTypes.arrayOf(PropTypes.object)
};

export default LinkMenu;
