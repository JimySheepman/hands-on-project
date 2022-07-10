import React, { Component } from 'react'
import { connect } from 'react-redux';
import {     UncontrolledDropdown,    DropdownToggle,    DropdownMenu,    NavItem,    NavLink,    DropdownItem,    Badge} from 'reactstrap';
import { bindActionCreators } from "redux"
import * as cartActions from "../../redux/actions/cartActions"
import {Link} from "react-router-dom"
import alertify from "alertifyjs"


class CartSummary extends Component {
    removeFromCart(product) {
        this.props.actions.removeFromCart(product);
        alertify.error(product.productName+ " delete to cart")
    }
    renderEmpty() {
        return (
            <NavItem>
                <NavLink>Cart is empty</NavLink>
            </NavItem>
        )
    }
    renderSummary() {
        return (
            <UncontrolledDropdown nav inNavbar>
                <DropdownToggle nav caret>
                    Carts
      </DropdownToggle>
                <DropdownMenu right>
                    {
                        this.props.cart.map(cartItem => (
                            <DropdownItem key={cartItem.product.id}>
                                <Badge color="danger" onClick={()=>this.removeFromCart(cartItem.product)}>-</Badge>
                                {cartItem.product.productName}
                                <Badge color="success">{cartItem.quantity}</Badge>
                            </DropdownItem>
                        ))
                    }
                    <DropdownItem divider />
                    <DropdownItem><Link to={"/cart"}>go to cart</Link>
                        
        </DropdownItem>
                </DropdownMenu>
            </UncontrolledDropdown>
        )
    }
    render() {
        return (
            <div>
                {
                    this.props.cart.length > 0 ? this.renderSummary() : this.renderEmpty()
                }

            </div>
        )
    }
}
function mapStateToProps(state) {
    return {
        cart: state.cartReducer
    }
}

function mapDispacthToProps(dispatch) {
    return {
        actions: {
            removeFromCart:bindActionCreators(cartActions.removeFromCart, dispatch)
        }
    }
}
export default connect(mapStateToProps,mapDispacthToProps)(CartSummary);