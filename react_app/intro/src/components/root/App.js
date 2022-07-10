import React, { Component } from 'react';
import { Route, Switch } from 'react-router-dom';
import { Container } from 'reactstrap';
import CartDetail from '../cart/CartDetail';
import Navi from '../navi/Navi';
import Dashboard from './Dashboard';


export default class App extends Component {

  render() {

    return (

      <Container>
        <Navi />
        <Switch>
          <Route path="/" exact component={Dashboard} />
          <Route path="/product" exact component={Dashboard} />
          <Route path="/cart" exact component={CartDetail} />
          <Route path="/" exact component={Dashboard} />
        </Switch>
      </Container>
    );

  }
}