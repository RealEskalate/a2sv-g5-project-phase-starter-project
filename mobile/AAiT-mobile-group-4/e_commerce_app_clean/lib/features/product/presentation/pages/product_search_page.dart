import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../injection_container.dart';
import '../bloc/product_bloc.dart';
import '../widgets/components/modal_sheet_widget.dart';
import '../widgets/components/product_card.dart';
import '../widgets/components/styles/text_style.dart';

class ProductSearchPage extends StatefulWidget {
  const ProductSearchPage({super.key});

  @override
  State<ProductSearchPage> createState() => _ProductSearchPageState();
}

class _ProductSearchPageState extends State<ProductSearchPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: BlocProvider(
        create: (_) => sl<ProductBloc>()..add(LoadAllProductEvent()),
        child: Container(
          padding: const EdgeInsets.all(32),
          child: Column(
            children: [
              Row(
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  IconButton(
                    onPressed: () {
                      Navigator.pop(context);
                    },
                    icon:  Icon(
                      Icons.arrow_back_ios_rounded,
                      color: Theme.of(context).primaryColor
                    ),
                  ),
                  const SizedBox(width: 60),
                  const CustomTextStyle(
                    name: 'Search Product',
                    weight: FontWeight.w500,
                    size: 16,
                  ),
                ],
              ),
              SingleChildScrollView(
                scrollDirection: Axis.horizontal,
                child: Row(
                  children: [
                   SizedBox(
                      width: 270,
                      height: 48,
                      child: Stack(
                        children: [
                          const TextField(
                            decoration: InputDecoration(
                              hintText: 'Leather',
                              enabledBorder: OutlineInputBorder(
                                borderSide: BorderSide(
                                    color: Color.fromRGBO(217, 217, 217, 1),
                                    width: 1.0),
                              ),
                              focusedBorder: OutlineInputBorder(
                                borderSide: BorderSide(
                                    color: Color.fromRGBO(217, 217, 217, 1),
                                    width: 1.0),
                              ),
                            ),
                          ),
                          Positioned(
                              left: 230,
                              child: IconButton(
                                onPressed: null,
                                icon: Icon(
                                  Icons.arrow_forward,
                                  color: Theme.of(context).primaryColor,
                                ),
                              )),
                        ],
                      ),
                    ),
                    const SizedBox(
                      width: 7,
                    ),
                    GestureDetector(
                      onTap: () {
                        showModalBottomSheet(
                          context: context,
                          builder: (BuildContext context) {
                            return const ModalSheetComponent();
                          },
                        );
                      },
                      child: Container(
                        width: 48,
                        height: 48,
                        decoration: BoxDecoration(
                          color: Theme.of(context).primaryColor,
                          borderRadius: BorderRadius.circular(8),
                        ),
                        child: const Icon(
                          Icons.filter_list_rounded,
                          color: Colors.white,
                        ),
                      ),
                    )
                  ],
                ),
              ),
              const SizedBox(
                height: 32,
              ),
              BlocBuilder<ProductBloc,ProductState>(builder: (context, state) {
                if (state is ProductLoading) {
                  return const Center(
                    child: CircularProgressIndicator(),
                  );
                } else if (state is LoadedAllProductState) {
                  return Expanded(
                      child: ListView.builder(
                          itemCount: state.products.length,
                          itemBuilder: (context, index) {
                            return MyCardBox(product: state.products[index]);
                          }));
                } else if (state is ProductErrorState) {
                  return Center(
                    child: Text(state.message),
                  );
                } else {
                  return const Center(child: Text('No products'),);
                }
              })
            ],
          ),
        ),
      ),
    );
  }
}
