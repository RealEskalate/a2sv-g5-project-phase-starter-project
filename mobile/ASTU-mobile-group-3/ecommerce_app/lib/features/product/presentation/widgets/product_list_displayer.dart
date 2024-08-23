import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../bloc/product_bloc.dart';
import '../bloc/product_events.dart';
import '../bloc/product_states.dart';
import '../pages/single_product_page.dart';
import 'product_card.dart';

class ProductListDisplayer extends StatelessWidget {
  const ProductListDisplayer({
    super.key,
  });
  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: BlocBuilder<ProductBloc, ProductStates>(
        builder: (context, state) {
          if (state is LoadedAllProductState) {
            return ListView.builder(
              itemCount: state.data.length,
              itemBuilder: (context, index) {
                return GestureDetector(
                  key: Key(state.data[index].id),
                  onTap: () {
                    BlocProvider.of<ProductBloc>(context).add(
                      GetSingleProductEvents(
                        id: state.data[index].id,
                      ),
                    );
                    Navigator.pushNamed(context, SingleProduct.routes);
                  },
                  child: ProductCard(
                    imageUrl: state.data[index].imageUrl,
                    price: state.data[index].price,
                    productName: state.data[index].name,
                    productType: state.data[index].description,
                    rating: '23',
                  ),
                );
              },
            );
          } else if (state is ErrorState) {
            return Center(
              child: Text(state.message),
            );
          } else if (state is LoadingState) {
            return const Center(child: CircularProgressIndicator());
          } else {
            return const Center(
              child: Text('No Data'),
            );
          }
        },
      ),
    );
  }
}
