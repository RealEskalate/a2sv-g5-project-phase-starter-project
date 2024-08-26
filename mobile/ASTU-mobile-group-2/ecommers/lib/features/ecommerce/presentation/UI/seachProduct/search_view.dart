import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../state/product_bloc/product_bloc.dart';
import '../../state/product_bloc/product_state.dart';
import '../home/product_image.dart';

class SearchView extends StatelessWidget {
  final String text;

  const SearchView({
    super.key,
    required this.text,
  
  });

  @override
  Widget build(BuildContext context) {
    List<dynamic> data = [];
    return BlocBuilder<ProductBloc, ProductState>(
      builder: (context, state) {
        if (state is LoadedAllProductState){
          data = state.products;
        }
        return ListView.builder(
            itemCount: data.length,
            itemBuilder: (context, index) {
              final product = data[index];
              
              if (product.name.toString().startsWith(text)) {
                return ProductImage(
                  disc: product.description,
                  imageUrl: product.imageUrl,
                  price: product.price,
                  title: product.name,
                  id: product.id,
                  senderId:  '',
                  senderName: '',
                );
              }
              return null;
            
            });
      },
    );
  }
}
