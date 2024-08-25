import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../core/themes/themes.dart';
import '../bloc/product_bloc.dart';
import '../bloc/product_events.dart';
import '../bloc/product_states.dart';
import '../widgets/product_widgets.dart';
import '../widgets/skeleton_loading.dart';

class SearchProduct extends StatefulWidget {
  static String routes = '/search_product';

  const SearchProduct({super.key});

  @override
  State<StatefulWidget> createState() => _SearchProduct();
}

class _SearchProduct extends State<SearchProduct> with AppBars {
  final TextEditingController searchControl = TextEditingController();
  final TextEditingController catagory = TextEditingController();
  bool searchReady = false;
  RangeValues theRange = const RangeValues(10, 90);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: normalAppBar('Search Product', () {
        BlocProvider.of<ProductBloc>(context).add(LoadAllProductEvents());
        Navigator.pop(context);
      }),
      body: Column(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Row(
            children: [
              Expanded(
                child: SearchInput(
                  hint: 'Search Item',
                  control: searchControl,
                  search: () {},
                  onChange: (text) {
                    ///BlocProvider.of<ProductBloc>(context).add(RefreshEvent());
                    setState(() {});
                  },
                ),
              ),
              Padding(
                padding: const EdgeInsets.only(right: 20),
                child: FilledButton(
                  style: FilledButton.styleFrom(
                      backgroundColor: MyTheme.ecBlue,
                      padding: const EdgeInsets.symmetric(vertical: 16),
                      shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(10))),
                  onPressed: () {
                    setState(() {
                      searchReady = !searchReady;
                    });
                  },
                  child: const Icon(Icons.filter_list),
                ),
              )
            ],
          ),
          BlocBuilder<ProductBloc, ProductStates>(
            builder: (context, state) {
              if (state is LoadedAllProductState) {
                return Expanded(
                  child: ListView.builder(
                    itemCount: state.data.length,
                    itemBuilder: (context, index) {
                      if (state.data[index].name
                          .toLowerCase()
                          .contains(searchControl.text.toLowerCase())) {
                        return ProductCard(
                          imageUrl: state.data[index].imageUrl,
                          productName: state.data[index].name,
                          price: state.data[index].price,
                          productType: state.data[index].description,
                          rating: '4.0',
                        );
                      } else {
                        return const SizedBox();
                      }
                    },
                  ),
                );
              } else {
                return ListView.builder(
                  itemCount: 10,
                  itemBuilder: (context, index) {
                    return const SkeletonLoading();
                  },
                );
              }
            },
          ),
          if (searchReady)
            Container(
              padding: const EdgeInsets.all(10),
              child: Column(
                children: [
                  CostumInput(
                    hint: '',
                    text: 'Catagory',
                    control: catagory,
                    fillColor: Colors.white,
                    borderColor: MyTheme.ecGrey,
                  ),
                  Row(
                    children: [
                      Expanded(
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            const Padding(
                              padding: EdgeInsets.symmetric(
                                  horizontal: 30, vertical: 10),
                              child: Text(
                                'Price',
                                style: TextStyle(fontWeight: FontWeight.bold),
                              ),
                            ),
                            Padding(
                              padding:
                                  const EdgeInsets.symmetric(horizontal: 10),
                              child: RangeSlider(
                                min: 0,
                                max: 100,
                                activeColor: MyTheme.ecBlue,
                                inactiveColor: MyTheme.ecGrey,
                                values: theRange,
                                divisions: 10,
                                onChanged: (newVal) {
                                  setState(() {
                                    theRange = newVal;
                                  });
                                },
                              ),
                            ),
                          ],
                        ),
                      )
                    ],
                  ),
                  Row(
                    children: [
                      Expanded(
                        child: Padding(
                            padding: const EdgeInsets.symmetric(horizontal: 10),
                            child:
                                FillCustomButton(press: () {}, label: 'Apply')),
                      ),
                    ],
                  ),
                ],
              ),
            ),
        ],
      ),
    );
  }
}
