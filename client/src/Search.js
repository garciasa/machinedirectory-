import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';

function Search({ onChange }) {
  function handleChange(e) {
    e.preventDefault();
    onChange(e.target.value);
  }

  return (
    <div className="search-box">
      <div className="sb-container">
        <span className="icon">
          <i className="fa fa-search" />
        </span>
        <input
          type="text"
          className="search"
          placeholder="search..."
          onChange={handleChange}
        />
        <Link to="/add">
          <span className="icon">
            <i className="fa fa-plus" />
          </span>
        </Link>
      </div>
    </div>
  );
}

Search.propTypes = {
  onChange: PropTypes.func.isRequired,
};

export default Search;
