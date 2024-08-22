import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:e_commerce_app/features/product/presentation/bloc/search/search_product_bloc.dart';
import 'package:e_commerce_app/features/product/presentation/bloc/update/update_product_bloc.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'widgets/widgets.dart';

class SearchPage extends StatelessWidget {
  SearchPage({super.key});

  @override
  Widget build(BuildContext context) {
    TextEditingController search_controller = TextEditingController();
    TextEditingController category_controller = TextEditingController();
    void update({required List<ProductEntity> products}) {
      context.read<SearchBloc>().add(
          ProductSearched(allProducts: products, name: search_controller.text));
    }

    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.white,
        centerTitle: true,
        automaticallyImplyLeading: false,
        leading: GoBack(),
        title: Text(
          "Search Product",
          style: TextStyle(
            fontSize: 16,
            fontWeight: FontWeight.w500,
          ),
        ),
      ),
      body: BlocBuilder<SearchBloc, SearchState>(
        builder: (context, searchstate) {
          List<ProductEntity> products = [];
          if (searchstate is SearchInitial || searchstate is SearchSuccess) {
            if (searchstate is SearchInitial) {
               products = searchstate.allProducts;
            } else if (searchstate is SearchSuccess) {
              products = searchstate.filtered;
            }

            List<Widget> allCards = products.map((Product) {
              return ItemCard(product: Product);
            }).toList();
            return Container(
              color: Colors.white,
              padding: EdgeInsets.all(20),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Row(
                    children: [
                      Expanded(
                        child: Container(
                          decoration: BoxDecoration(
                            border: Border.all(
                              color: Color.fromARGB(255, 217, 217, 217),
                              width: 1,
                            ),
                            borderRadius: BorderRadius.all(
                              Radius.circular(10),
                            ),
                            color: Colors.white,
                          ),
                          child: TextField(
                            maxLines: 1,
                            controller: search_controller,
                            decoration: InputDecoration(
                              border: InputBorder.none,
                              suffixIcon: IconButton(
                                icon: Icon(Icons.arrow_forward),
                                color: Color.fromARGB(255, 63, 81, 243),
                                onPressed: () {
                                  update(products: products);
                                },
                              ),
                              hintText: "Leather",
                              hintStyle: TextStyle(
                                fontSize: 20,
                                fontWeight: FontWeight.w400,
                              ),
                              contentPadding:
                                  EdgeInsets.only(left: 16, top: 10),
                            ),
                          ),
                        ),
                      ),
                      SizedBox(
                        width: 10,
                      ),
                      ButtonIcon(
                        buildcontext: context,
                        callback: () {
                          showModalBottomSheet(
                              context: context,
                              builder: (BuildContext context) {
                                return Container(
                                    color: Colors.white,
                                    padding: EdgeInsets.all(20),
                                    child: FilterBottom(
                                        category_controller:
                                            category_controller));
                              });
                        },
                        icon: Icons.filter_list,
                      )
                    ],
                  ),
                  SizedBox(
                    height: 20,
                  ),
                  // card
                  Expanded(
                    child: SizedBox(
                      height: 100,
                      child: ListView(
                        children: allCards,
                      ),
                    ),
                  ),
                  SizedBox(
                    height: 20,
                  ),
                ],
              ),
            );
          }

          ///
          else {
            return Center(
              child: Text("NOT found roducts"),
            );
          }
        },
      ),
    );
  }

  static _textField1(String title, [int lines = 1]) {
    return SizedBox(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            title,
            style: TextStyle(
              fontSize: 20,
            ),
          ),
          SizedBox(
            height: 10,
          ),
          Container(
            decoration: BoxDecoration(
              border: Border.all(color: Colors.grey),
              borderRadius: BorderRadius.all(
                Radius.circular(10),
              ),
              color: Color.fromARGB(255, 246, 241, 241),
            ),
            child: TextField(
              maxLines: lines,
              decoration: InputDecoration(
                border: InputBorder.none,
                // suffixIcon: Icon(Icons.pedal_bike),
              ),
            ),
          ),
        ],
      ),
    );
  }
}

class FilterBottom extends StatelessWidget {
  const FilterBottom({
    super.key,
    required this.category_controller,
  });

  final TextEditingController category_controller;

  @override
  Widget build(BuildContext context) {
    return Column(
      // height: 200,
      children: [
        TextFieldTitle(
            controller: category_controller,
            title: "Category",
            color: Colors.white,
            border: true,
            fontsize: 16),

        // ItemCard()
        SizedBox(height: 15),
        Text(
          'Price',
          style: TextStyle(
            fontSize: 16,
            fontWeight: FontWeight.w400,
          ),
        ),

        RangeSlider(
          activeColor: Colors.blue,
          values: RangeValues(2, 9),
          max: 10,
          min: 1,
          onChanged: (RangeValues newValue) {},
        ),
        SizedBox(
          height: 25,
        ),
        BackgroundButton(
          title: "APPLY",
        ),
      ],
    );
  }
}
