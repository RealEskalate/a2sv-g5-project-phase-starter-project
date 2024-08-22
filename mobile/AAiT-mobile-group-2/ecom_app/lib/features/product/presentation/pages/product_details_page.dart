import 'package:cached_network_image/cached_network_image.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../../domain/entities/product.dart';
import '../bloc/product_bloc.dart';
import '../widgets/chips.dart';
import '../widgets/custom_outlined_button.dart';

class ProductDetailsPage extends StatelessWidget {
  final Product product;
  const ProductDetailsPage({
    super.key,
    required this.product,
  });

  @override
  Widget build(BuildContext context) {
    List<int> shoeSize = [39, 40, 41, 42, 43, 44, 45];
    return Scaffold(
      // appBar: AppBar(),
      backgroundColor: const Color.fromRGBO(254, 254, 254, 1),
      body: SingleChildScrollView(
        child: BlocListener<ProductBloc, ProductState>(
          listener: (context, state) {
            if (state is ProductErrorState) {
              ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
                  content: Text('Failed to delete the product'), backgroundColor: Colors.red,));
            } else if (state is ProductDeletedState) {
              if (!context.mounted) return;
              ScaffoldMessenger.of(context).showSnackBar( SnackBar(
                  content: const Text('Product Deleted Successfully'), backgroundColor: Theme.of(context).primaryColor,));
              Navigator.of(context).pushNamed('/home');
            }
          },
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Stack(children: [
                CachedNetworkImage(
                  imageUrl: product.imageUrl,
                  width: double.infinity,
                  height: MediaQuery.of(context).size.height * 0.32,
                  fit: BoxFit.cover,
                ),
                Positioned(
                  top: 40,
                  left: 10,
                  child: Container(
                    width: 40.0,
                    height: 40.0,
                    decoration: BoxDecoration(
                        color: Colors.white,
                        borderRadius: BorderRadius.circular(20)),
                    child: Center(
                      child: IconButton(
                        onPressed: () {
                          Navigator.pop(context);
                        },
                        icon: Icon(
                          Icons.arrow_back_ios_new_rounded,
                          color: Colors.indigoAccent.shade400,
                        ),
                      ),
                    ),
                  ),
                ),
              ]),
              Padding(
                padding: const EdgeInsets.all(20),
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.start,
                  crossAxisAlignment: CrossAxisAlignment.start,
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        const Text(
                          "Chuck Taylor's",
                          style: const TextStyle(color: Colors.black45),
                        ),
                        Row(
                          children: [
                            Icon(
                              Icons.star_rate,
                              color: Colors.amber[400],
                            ),
                            const Text(
                              '(4.0)',
                              style: const TextStyle(color: Colors.black45),
                            )
                          ],
                        )
                      ],
                    ),
                    const SizedBox(
                      height: 10,
                    ),
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Text(
                          product.name,
                          style: const TextStyle(
                              fontWeight: FontWeight.w600, fontSize: 20),
                        ),
                        Text(
                          '\$${product.price}',
                          style: const TextStyle(fontWeight: FontWeight.w500),
                        )
                      ],
                    ),
                    const SizedBox(
                      height: 20,
                    ),
                    const Text(
                      'Size:',
                      style:
                          TextStyle(fontWeight: FontWeight.w600, fontSize: 16),
                    ),
                    const SizedBox(
                      height: 10,
                    ),
                    SizedBox(
                        height: 60,
                        child: ListView.separated(
                          physics: const AlwaysScrollableScrollPhysics(),
                          scrollDirection: Axis.horizontal,
                          itemCount: shoeSize.length,
                          itemBuilder: (BuildContext context, int index) {
                            if (39 == (shoeSize[index]) ||
                                41 == (shoeSize[index])) {
                              return Chips(
                                number: shoeSize[index],
                                selected: true,
                              );
                            } else {
                              return Chips(
                                number: shoeSize[index],
                                selected: false,
                              );
                            }
                          },
                          separatorBuilder: (BuildContext context, _) {
                            return const SizedBox(
                              width: 6,
                            );
                          },
                        )),
                    const SizedBox(
                      height: 20,
                    ),
                    Text(product.description),
                    const SizedBox(
                      height: 60,
                    ),
                    BlocBuilder<ProductBloc, ProductState>(
                      builder: (context, state) {
                        return Row(
                          mainAxisAlignment: MainAxisAlignment.spaceAround,
                          children: [
                            CustomOutlinedButton(
                                backgroundColor: Colors.white,
                                foregroundColor:
                                    const Color.fromRGBO(255, 19, 19, 0.79),
                                borderColor:
                                    const Color.fromRGBO(255, 19, 19, 0.79),
                                buttonWidth: 120,
                                buttonHeight: 45,
                                child: state is ProductLoading
                                    ? const CircularProgressIndicator(
                                        valueColor: AlwaysStoppedAnimation<
                                                Color>(
                                            Color.fromRGBO(255, 19, 19, 0.79)),
                                      )
                                    : const Text(
                                        'DELETE',
                                        style: const TextStyle(
                                            fontWeight: FontWeight.w600),
                                      ),
                                onPressed: () => {
                                      context.read<ProductBloc>().add(
                                          DeleteProductEvent(id: product.id))
                                    }),
                            CustomOutlinedButton(
                                backgroundColor: Theme.of(context).primaryColor,
                                foregroundColor: Colors.white,
                                borderColor: Theme.of(context).primaryColor,
                                buttonWidth: 120,
                                buttonHeight: 45,
                                child: const Text(
                                  'UPDATE',
                                  style: const TextStyle(
                                      fontWeight: FontWeight.w600),
                                ),
                                onPressed: () {
                                  Navigator.of(context).pushNamed('/update',
                                      arguments: Product(
                                          id: product.id,
                                          imageUrl: product.imageUrl,
                                          name: product.name,
                                          price: product.price,
                                          description: product.description));
                                })
                          ],
                        );
                      },
                    )
                  ],
                ),
              )
            ],
          ),
        ),
      ),
    );
  }
}
