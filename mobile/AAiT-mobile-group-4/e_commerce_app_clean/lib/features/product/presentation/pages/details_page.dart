import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../domain/entities/product_entity.dart';
import '../bloc/product_bloc.dart';
import '../widgets/components/size_cards.dart';
import '../widgets/components/styles/custom_button.dart';
import '../widgets/components/styles/text_style.dart';

class DetailsPage extends StatelessWidget {
  final ProductEntity selectedProduct;

  const DetailsPage({
    super.key,
    required this.selectedProduct,
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SingleChildScrollView(
        child: BlocListener<ProductBloc, ProductState>(
          listener: (context, state) {
            if (state is ProductDeletedState) {
              ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
                  content: Text('successfully deleted product')));
              Navigator.pushNamed(context, '/home_page');
            } else if (state is ProductErrorState) {
              ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
                  key: Key('snackbar_error'), content: Text('error')));
            }
          },
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Stack(
                children: [
                  SizedBox(
                    width: MediaQuery.of(context).size.width,
                    height: 286,
                    child: Image.network(
                      selectedProduct.imageUrl,
                      fit: BoxFit.cover,
                    ),
                  ),
                  Positioned(
                    left: 24,
                    top: 25,
                    child: IconButton(
                      onPressed: () {
                        Navigator.pop(context);
                      },
                      icon: Icon(
                        Icons.arrow_back,
                        color: Theme.of(context).primaryColor,
                      ),
                      style: IconButton.styleFrom(
                        backgroundColor: Colors.white,
                        shape: const CircleBorder(),
                      ),
                    ),
                  )
                ],
              ),
              const SizedBox(height: 21),
              Container(
                padding: const EdgeInsets.symmetric(horizontal: 32),
                child: const Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    CustomTextStyle(
                      name: 'Men\'s shoe',
                      weight: FontWeight.w400,
                      size: 16,
                      color: Color.fromRGBO(170, 170, 170, 1.0),
                    ),
                    Row(
                      children: [
                        Icon(
                          Icons.star,
                          color: Color.fromRGBO(255, 215, 0, 1),
                        ),
                        CustomTextStyle(
                          name: '4',
                          weight: FontWeight.w400,
                          size: 16,
                          color: Color.fromRGBO(170, 170, 170, 1.0),
                          family: 'Sora',
                        ),
                      ],
                    ),
                  ],
                ),
              ),
              const SizedBox(height: 16),
              Container(
                padding: const EdgeInsets.symmetric(horizontal: 32),
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    CustomTextStyle(
                      name: selectedProduct.name,
                      weight: FontWeight.w600,
                      size: 24,
                    ),
                    CustomTextStyle(
                      name: '\$${selectedProduct.price}',
                      weight: FontWeight.w500,
                      size: 16,
                    ),
                  ],
                ),
              ),
              const SizedBox(height: 20),
              const Padding(
                padding: EdgeInsets.only(left: 32),
                child: CustomTextStyle(
                  name: 'Size: ',
                  weight: FontWeight.w500,
                  size: 20,
                ),
              ),
              SizedBox(
                width: 500,
                height: 60,
                child: ListView.builder(
                  itemCount: 6,
                  scrollDirection: Axis.horizontal,
                  itemBuilder: (context, index) {
                    return SizeCards(
                      value: index + 39 == 41 ? true : false,
                      size: index + 39,
                    );
                  },
                ),
              ),
              const SizedBox(height: 32),
              SingleChildScrollView(
                child: Container(
                  padding: const EdgeInsets.only(left: 10),
                  height: 240,
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        selectedProduct.description,
                        style: const TextStyle(
                          fontFamily: 'Poppins',
                          fontSize: 14,
                          fontWeight: FontWeight.w500,
                        ),
                      ),
                      
                      Row(
                        mainAxisAlignment: MainAxisAlignment.spaceAround,
                        children: [
                          CustomButton(
                            pressed: () {
                              context.read<ProductBloc>().add(
                                  DeleteProductEvent(id: selectedProduct.id));
                              //Navigator.of(context).pop('delete');
                            },
                            name: 'DELETE',
                            width: 152,
                            height: 50,
                            fgcolor: Theme.of(context).secondaryHeaderColor,
                            textBgColor: Theme.of(context).secondaryHeaderColor,
                            bgcolor: Colors.white,
                          ),
                          const SizedBox(
                            width: 16,
                          ),
                          CustomButton(
                            pressed: () {
                              Navigator.pushNamed(context, '/update_page',
                                  arguments: selectedProduct);
                            },
                            name: 'UPDATE',
                            width: 152,
                            height: 50,
                            textBgColor: Colors.white,
                            fgcolor: Colors.white,
                            bgcolor: Theme.of(context).primaryColor,
                          ),
                        ],
                      ),
                    ],
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
