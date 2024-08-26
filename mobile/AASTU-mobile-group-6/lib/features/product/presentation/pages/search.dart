import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_event.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/search/search_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/search/search_event.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/search/search_state.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/widgets/bottomnavbar.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/get_user/get_user_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/get_user/get_user_event.dart';
import 'package:ecommerce_app_ca_tdd/locator.dart';
import "package:flutter/material.dart";
import 'package:ecommerce_app_ca_tdd/extra/overflow_card.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';
import 'package:ecommerce_app_ca_tdd/extra/search_func.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/home.dart';
import 'package:ecommerce_app_ca_tdd/extra/resusetext.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/details.dart';
import 'package:ecommerce_app_ca_tdd/models/product.dart';
import 'package:ecommerce_app_ca_tdd/models/ext_product.dart';

import '../bloc/home_state.dart';

class searchPage extends StatefulWidget {
  const searchPage({super.key});

  @override
  State<searchPage> createState() => _searchPageState();
}

class _searchPageState extends State<searchPage> {
  TextEditingController search_term = TextEditingController();

  Future<void> _refresh() {
    context.read<SearchBloc>().add(LoadAllProductEvent());

    return Future.delayed(Duration(seconds: 3));
  }

  void _performSearch() {
    final search = search_term.text;
    context.read<SearchBloc>().add(SearchProductEvent(search));
  }

  @override
  Widget build(BuildContext context) {
    BlocProvider(
      create: (context) => sl.get<SearchBloc>()..add(LoadAllProductEvent()),
    );

    return Scaffold(
      bottomNavigationBar: Container(child: Bottomnavbar()),

      appBar: AppBar(
        automaticallyImplyLeading: false,
        title: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            IconButton(
                onPressed: () {
                  Navigator.pushNamed(context, '/home');
                },
                icon: Icon(
                  Icons.arrow_back_ios_new,
                  color: Color.fromARGB(255, 63, 81, 243),
                  size: 20,
                )),
             Center(
              child: Text("Search Product",
              style: GoogleFonts.poppins(),
              ),
            ),
            const SizedBox(
              height: 60,
              width: 60,
            )
          ],
        ),
      ),
      body: SingleChildScrollView(
        child: Center(
          child: Container(
            margin: EdgeInsets.only(top: 10),
            width: MediaQuery.of(context).size.width *
                0.9, //---------------OVERFLOW!!

            child: Column(
              children: [
                Container(
                  child: Row(
                    children: [
                      Container(
                        margin: EdgeInsets.only(left: 7),
                        child: SizedBox(
                          width: MediaQuery.of(context).size.width *
                              0.70, 
                          height: 48,
                          child: Container(
                            decoration: BoxDecoration(
                                border: Border.all(
                                    color: Color.fromRGBO(217, 217, 217, 1)),
                                borderRadius:
                                    BorderRadius.all(Radius.circular(8))),
                            child: Expanded(
                              child: TextField(
                                controller: search_term,
                                decoration: InputDecoration(
                                  suffixIcon: IconButton(
                                      onPressed: () {
                                        final search = search_term.text;
                                        context
                                            .read<SearchBloc>()
                                            .add(SearchProductEvent(search));
                                      },
                                      icon: Icon(
                                        Icons.arrow_forward,
                                        color: Color.fromARGB(255, 63, 81, 243),
                                      )),
                                  border: InputBorder.none,
                                  hintText: "Leather",
                                  hintStyle: TextStyle(color: Colors.grey),
                                  contentPadding: EdgeInsets.all(10)
                                ),
                              ),
                            ),
                          ),
                        ),
                      ),

                      SizedBox(
                        width: 15,
                      ),

                      Container(
  decoration: BoxDecoration(
    color: Color(0xff3F51F3),
    borderRadius: BorderRadius.circular(7), 
  ),
  child: Container(
    decoration: BoxDecoration(
      border: Border.all(
        width: 1,
        color: Color(0xff3F51F3),
      ),
      borderRadius: BorderRadius.circular(10), // Same border radius for inner container
    ),
    child: IconButton(
      onPressed: () {
        showModalBottomSheet(
          context: context,
          builder: (BuildContext context) {
            return const SizedBox(
              height: 338,
              child: about_product(),
            );
          },
        );
      },
      icon: Icon(
        Icons.filter_list,
        color: Colors.white,
      ),
    ),
  ),
),
],
                  ),
                ),

                SizedBox(height: 31),

                BlocBuilder<SearchBloc, SearchState>(
                  builder: (context, state) {
                    if (state is LoadingState) {
                      return const Center(
                        child: CircularProgressIndicator(),
                      );
                    } else if (state is LoadedState) {
                      return RefreshIndicator(
                        onRefresh: _refresh,
                        child: ListView.builder(
                          shrinkWrap: true,
                          physics: const NeverScrollableScrollPhysics(),
                          itemCount: state.data.length,
                          itemBuilder: (context, index) {
                            final product = state.data[index];
                            return InkWell(
                              onTap: () {
                                Navigator.push(
                                  context,
                                  MaterialPageRoute(
                                    builder: (context) =>
                                        DetailsPage(item: product),
                                  ),
                                );
                              },
                              child: OverflowCard(
                                product: product,
                              ),
                            );
                          },
                        ),
                      );
                    } else if (state is FailedState) {
                      return Center(
                        child: Text(state.message),
                      );
                    } else {
                      return const SizedBox();
                    }
                  },
                ),

                //
              ],
            ),
          ),
        ),
      ),
    );
  }
}
