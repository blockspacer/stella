import React, { useState } from "react";
import RoundButton from "../RoundButton/RoundButton";
import styled from "styled-components";
import CircleButton from "../CircleButton/CircleButton";
import { ReactComponent as FaceProfileIcon } from "../../img/face-profile.svg";
import { Typography, Card } from "@material-ui/core";
import { Book } from "proto/booksvc_pb";
import { User } from "proto/usersvc_pb";
import AppHeader from "../../AppHeader";

const Header = styled.div`
  display: flex;
  min-height: 84px;
  padding-left: 32px;
  padding-right: 32px;
  align-items: center;
  justify-content: space-between;
`;

const SearchForm = styled.form`
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const SearchContainer = styled.div`
  display: flex;
  min-height: 652px;
  flex-direction: column;
  align-items: center;
  justify-content: space-around;
`;

const SearchInput = styled.input`
  width: 750px;
  font-size: 96px;
  font-weight: 300;
  border: none;
`;

interface IndexPageProps {
  currentUserLoader?: () => [boolean, User | undefined]
  search?: (query: string) => [boolean, Book[]]
}

const defaultSearch = (q: string) : [boolean, Book[]] => ([false, []])
const defaultCurrentUserLoader = () : [boolean, User | undefined] => ([false, undefined])

const IndexPage: React.SFC<IndexPageProps> = ({
  search = defaultSearch,
  currentUserLoader = defaultCurrentUserLoader
}) => {
  const [haveSearched, setSearched] = useState(false);
  const [query, setQuery] = useState('');
  if (!search) return null;
  const [loading, books] = search(query)
  const [loadingCurrentUser, currentUser] = currentUserLoader();
  return (
    <div>
      <AppHeader currentUser={currentUser} />
      <SearchContainer>
        <SearchForm
          onSubmit={e => {
            e.preventDefault();
            const target = e.target as HTMLFormElement;
            const input = target.elements.namedItem('search') as HTMLInputElement;
            input && setQuery(input.value);
            setSearched(true);
          }}
        >
          <SearchInput name="search" placeholder="Search for a book" />
          <Typography variant="title">
            Try: #All, #English or #popular to filter
          </Typography>
          {loading && <Typography>Loading</Typography>}
          <div>
            {books.map(b => (
              <Card key={b.getId()}>
                <img src={b.getCoverImage()} title={b.getName()} />
                <Typography>{b.getName()} - {b.getAuthorsList().join(',')}</Typography>
              </Card>
            ))}
          </div>
        </SearchForm>
      </SearchContainer>
    </div>
  );
};

export default IndexPage;
