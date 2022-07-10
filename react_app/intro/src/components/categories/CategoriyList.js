import React, { Component } from 'react'
import { connect } from "react-redux"
import { ListGroup, ListGroupItem, Badge } from 'reactstrap'
import { bindActionCreators } from "redux"
import * as categoryActions from "../../redux/actions/categoryActions"
import * as productActions from "../../redux/actions/productActions"

class CategoriyList extends Component {
    componentDidMount() {
        this.props.actions.getCategories()
    }

    selectCategor = (category) => {
        this.props.actions.changeCategory(category)
        this.props.actions.getProducts(category.id)
    }
    render() {
        return (
            <div>
                <h3>
                    <Badge color="warning">
                        Categories
                    </Badge></h3>
                <ListGroup>
                    {
                        this.props.categories.map(category => (
                            <ListGroupItem active={category.id === this.props.currentCategory.id} onClick={() => this.selectCategor(category)} key={category.id}>
                                {category.categoryName}
                            </ListGroupItem>
                        ))}
                </ListGroup>
            </div>
        )
    }
}
function mapStateToProps(state) {
    return {
        currentCategory: state.changeCategoryReducer,
        categories: state.categoryListReducer
    }
}

function mapDispacthToProps(dispatch) {
    return {
        actions: {
            getCategories: bindActionCreators(categoryActions.getCategories, dispatch),
            getProducts: bindActionCreators(productActions.getProducts, dispatch),
            changeCategory: bindActionCreators(categoryActions.changeCategory, dispatch)
        }
    }
}
export default connect(mapStateToProps, mapDispacthToProps)(CategoriyList)