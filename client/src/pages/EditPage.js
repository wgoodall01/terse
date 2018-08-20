import React from 'react';
import gql from 'graphql-tag';
import {Query} from 'react-apollo';
import './EditPage.css';

import LinkMenu from 'components/LinkMenu.js';
import LinkEdit from 'components/LinkEdit.js';

const getFromShort = (links, short) => links.filter(e => e.short === short).shift();

const EditPage = ({
  match: {
    params: {slug: currentShort}
  },
  history
}) => (
  <Query query={linksQuery}>
    {({loading, error, data}) => {
      if (loading) return <h1>Loading...</h1>;
      if (error) return <h1>Error: {error.message}</h1>;

      const currentLink = getFromShort(data.link, currentShort) || {};

      return (
        <div className="EditPage">
          <div className="EditPage_menu">
            <LinkMenu links={data.link} current={(currentLink || {id: null}).id} />
          </div>
          <div className="EditPage_detail">
            <LinkEdit key={currentLink.id} link={currentLink} onSubmit={() => {}} />
          </div>
        </div>
      );
    }}
  </Query>
);

const linksQuery = gql`
  query {
    link {
      id
      long
      long
      short
    }
  }
`;

export default EditPage;
