// Code generated by qtc from "organizers-companies.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package v1

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func StreamOrganizersCompanies(qw422016 *qt422016.Writer,
	organizerFeature OrganizerFeature,
	headerProfiles []SocialProviderUser,
	companies []Company,
	ukrainianUniversities []University,
	czechUniversities []University,
	userCompanyFavoriteMap map[int64]bool,
	authQueryParams string,
) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">

<head>
	<title>Companies using `)
	qw422016.E().S(organizerFeature.Organizer.Title)
	qw422016.N().S(` | ReadyToTouch</title>
	<meta name="title" content="Companies using `)
	qw422016.E().S(organizerFeature.Organizer.Title)
	qw422016.N().S(` | ReadyToTouch">
	<meta name="description" content="`)
	qw422016.E().S(organizerFeature.Organizer.Description)
	qw422016.N().S(`">
    <meta name="keywords" content="`)
	qw422016.E().S(organizerFeature.Organizer.Title)
	qw422016.N().S(`">

	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    `)
	streamfavicon(qw422016)
	qw422016.N().S(`
    `)
	streamorganizersFonts(qw422016)
	qw422016.N().S(`
    `)
	streamorganizersCompaniesStyles(qw422016)
	qw422016.N().S(`
</head>

<body class="organizer-companies-inner">
<main class="main-wrapper">
	<header class="header">
  <div class="header__wrapper">
    <a href="/organizers" class="header__logo">
      <img
        width="129"
        height="32"
        class="header__logo-img"
        src="/assets/images/pages/organizer/`)
	qw422016.E().S(organizerFeature.Organizer.Logo)
	qw422016.N().S(`"
        alt="organizer logo"
      />
    </a>
    `)
	var navigation = toFeatureNavigation(organizerFeature.Path)

	qw422016.N().S(`
    <ul class="header__nav">
      <li class="header__nav-item">
        <a href="`)
	qw422016.E().S(navigation.companiesURL)
	qw422016.N().S(`" class="header__nav-link `)
	qw422016.E().S(navigation.companiesActive)
	qw422016.N().S(`">Companies</a>
      </li>
      <li class="header__nav-item">
        <a href="`)
	qw422016.E().S(navigation.vacanciesURL)
	qw422016.N().S(`" class="header__nav-link `)
	qw422016.E().S(navigation.vacanciesActive)
	qw422016.N().S(`">Vacancies</a>
      </li>
    </ul>
    `)
	streamorganizersHeaderStars(qw422016)
	qw422016.N().S(`
    `)
	if len(headerProfiles) > 0 {
		qw422016.N().S(`
    `)
		streamorganizersHeaderProfile(qw422016, headerProfiles)
		qw422016.N().S(`
    `)
	} else {
		qw422016.N().S(`
    <a href="/organizers/`)
		qw422016.E().S(organizerFeature.Organizer.Alias)
		qw422016.N().S(`/welcome`)
		qw422016.E().S(authQueryParams)
		qw422016.N().S(`" class="button button--small-padding button--black header__login-button">Log in</a>
    `)
	}
	qw422016.N().S(`
  </div>
</header>

<div class="container">
  <nav aria-label="breadcrumb" aria-labelledby="navigation through the bread crumbs" class="breadcrumb">
    <ul class="breadcrumb__list">
      <li class="breadcrumb__item">
        <a class="breadcrumb__link" href="/">Main</a>
      </li>
      <li class="breadcrumb__item">
        <a class="breadcrumb__link" href="/organizers">Organizers</a>
      </li>
      <li class="breadcrumb__item">
        <a class="breadcrumb__link" href="javascript:void(0);">`)
	qw422016.E().S(organizerFeature.Organizer.Title)
	qw422016.N().S(`</a>
      </li>
      <li class="breadcrumb__item">
        <span class="breadcrumb__page" aria-current="page">`)
	qw422016.E().S(organizerFeature.Title)
	qw422016.N().S(`</span>
      </li>
    </ul>
  </nav>
</div>

<section class="search-container container">
  <div class="search search--projects search--organizer">
    <div class="search__input-group">
      <input class="search__input" id="js-company-query" type="search" name="search" placeholder="Search" list="js-company-query-datalist" />
      <datalist id="js-company-query-datalist">
        `)
	for _, company := range companies {
		qw422016.N().S(`
          <option value="`)
		qw422016.E().S(company.Name)
		qw422016.N().S(`"></option>
        `)
	}
	qw422016.N().S(`
      </datalist>
      <img class="search__icon" alt="Search icon" width="20" height="20" src="/assets/images/pages/common/search.svg" />
    </div>
  </div>
</section>

<div class="search-result mt-32">
  <div class="container">
    <div class="search-result__wrapper">
      <!-- filters -->
      <aside class="search-result__filters">
  <div class="search-result__filter-group search-result__filter-group--wide">
    <div class="search-result__filter-header">
      <h2 class="search-result__filter-headline">Filters:</h2>
      <button id="js-criteria-reset" class="button button--light-link search-result__filter-headline-reset" style="visibility: hidden;">Reset all</button>
    </div>

    <div class="filters">
      <!-- Company type -->
      <div class="filters__group">
        <header class="filters__header">
          <h4 class="filters__headline">Company type</h4>
        </header>
        <div class="filters__elements">
          <label class="checkbox filters__element">
            <input class="js-criteria-company-type checkbox__input" type="checkbox" data-alias="product" />
            <span class="checkbox__element"></span>
            Product
          </label>
          <label class="checkbox filters__element">
            <input class="js-criteria-company-type checkbox__input" type="checkbox" data-alias="startup" />
            <span class="checkbox__element"></span>
            Startup
          </label>
        </div>
      </div>

      <div class="filters__group">
        <header class="filters__header">
          <h4 class="filters__headline">Industry</h4>
        </header>
        <div class="filters__elements">
          `)
	for _, industry := range industries {
		qw422016.N().S(`
          <label class="checkbox filters__element">
            <input class="js-criteria-company-industry checkbox__input" type="checkbox" data-alias="`)
		qw422016.E().S(industry.Alias)
		qw422016.N().S(`" />
            <span class="checkbox__element"></span>
            `)
		qw422016.E().S(industry.Name)
		qw422016.N().S(`
          </label>
          `)
	}
	qw422016.N().S(`
        </div>
      </div>

      <!-- Other -->
      <div class="filters__group">
        <header class="filters__header">
          <h4 class="filters__headline">Other</h4>
        </header>
        <div class="filters__elements">
          <label class="checkbox filters__element">
            <input class="js-criteria-has-employees-from-country checkbox__input" type="checkbox" data-alias="ukraine" />
            <span class="checkbox__element"></span>
            Has Ukrainian employees
            <img
              class="checkbox__content-image"
              alt="Flag of Ukraine with waves"
              width="24"
              height="24"
              src="/assets/images/pages/online-new/ua_flag_with_waves.svg"
            />
          </label>
          <label class="checkbox filters__element">
            <input class="js-criteria-has-employees-from-country checkbox__input" type="checkbox" data-alias="czechia" />
            <span class="checkbox__element"></span>
            Has Czechs employees
            <img
              class="checkbox__content-image"
              alt="Flag of Czechia"
              width="24"
              height="24"
              src="/assets/images/pages/online-new/cz_flag.svg"
            />
          </label>
          <label class="checkbox filters__element">
            <input id="js-criteria-in-favorites" class="checkbox__input" type="checkbox" />
            <span class="checkbox__element"></span>
            Favorites
          </label>
          `)
	qw422016.N().S(`
        </div>
      </div>

      <!-- /Other -->

    </div>
  </div>
</aside>

      <!-- /filters -->
      <div class="search-result--group">
        <!-- selected filters -->
        <div class="filter-used filter-used--small" style="visibility: hidden;">
          <div class="filter-used__title">Applied filters:</div>
          <ul id="js-company-selected-criteria" class="filter-used__list"></ul>
        </div>
        <!-- /selected filters -->

        <div id="search_result_list" class="search-result__list">
          <p class="search-result-found"><span id="js-result-count" class="search-result-found__amount">`)
	qw422016.N().D(len(companies))
	qw422016.N().S(`</span> results</p>
          <!-- card list -->
          <div class="search-result__cards row-gap-8 mt-24">
            `)
	for _, company := range companies {
		qw422016.N().S(`
            <div class="js-company card"
                 data-company-id="`)
		qw422016.N().DL(company.ID)
		qw422016.N().S(`"
                 data-company-name="`)
		qw422016.E().S(company.Name)
		qw422016.N().S(`"
                 data-company-type="`)
		qw422016.E().S(string(company.Type))
		qw422016.N().S(`"
                 data-company-industries="`)
		qw422016.E().S(aliases(company.Industries))
		qw422016.N().S(`"
                 data-company-has-employees-from-countries="`)
		qw422016.E().S(aliases(company.HasEmployeesFromCountries))
		qw422016.N().S(`"
            >
              <aside class="card__action">
                `)
		if userCompanyFavoriteMap[company.ID] {
			qw422016.N().S(`
                  <button class="js-company-favorite favorite card__action-button button-group__item in-favorite" title="Remove from favorites">
                    <svg xmlns="http://www.w3.org/2000/svg" xml:space="preserve" class="favorite__icon" viewBox="0 0 28 28">
                      <path
                        d="m14.5 22.1-.5-.3-.5.3-6.8 4.2c-.5.3-1.1-.1-.9-.7L7.5 18l.1-.6-.4-.4-5.9-5.2c-.3-.3-.3-.6-.2-.8.1-.2.3-.4.5-.4l7.9-.7.6-.1.2-.6 2.9-7.4c.2-.5 1-.5 1.2 0l3.1 7.3.2.5.6.1 7.9.7c.2 0 .4.2.5.5.1.3 0 .6-.2.7l-5.9 5.2-.4.4.1.6 1.8 7.7c.1.3 0 .5-.2.6-.2.1-.5.2-.8 0l-6.6-4z"
                      />
                    </svg>
                  </button>
                `)
		} else {
			qw422016.N().S(`
                  <button class="js-company-favorite favorite card__action-button button-group__item" title="Add to favorite">
                    <svg xmlns="http://www.w3.org/2000/svg" xml:space="preserve" class="favorite__icon" viewBox="0 0 28 28">
                      <path
                        d="m14.5 22.1-.5-.3-.5.3-6.8 4.2c-.5.3-1.1-.1-.9-.7L7.5 18l.1-.6-.4-.4-5.9-5.2c-.3-.3-.3-.6-.2-.8.1-.2.3-.4.5-.4l7.9-.7.6-.1.2-.6 2.9-7.4c.2-.5 1-.5 1.2 0l3.1 7.3.2.5.6.1 7.9.7c.2 0 .4.2.5.5.1.3 0 .6-.2.7l-5.9 5.2-.4.4.1.6 1.8 7.7c.1.3 0 .5-.2.6-.2.1-.5.2-.8 0l-6.6-4z"
                      />
                    </svg>
                  </button>
                `)
		}
		qw422016.N().S(`

                <a href="`)
		qw422016.E().S(organizerFeature.Path)
		qw422016.N().S(`/`)
		qw422016.E().S(company.LinkedInProfile.Alias)
		qw422016.N().S(`/" class="button-group__item" title="View statistics">
                  <img width="20" height="20" alt="icon stats" src="/assets/images/pages/common/stats.svg" />
                </a>
              </aside>
              <figure class="card__header">
                <div class="card__image-overlay card__image-overlay--small">
                  <img
                    width="18"
                    height="18"
                    class="card__image card__image--preview"
                    alt="card image preview icon"
                    src="/assets/images/pages/common/image-preview.svg"
                  />
                </div>
                <figcaption class="card__header-caption">
                  <a href="`)
		qw422016.E().S(organizerFeature.Path)
		qw422016.N().S(`/`)
		qw422016.E().S(company.LinkedInProfile.Alias)
		qw422016.N().S(`/" class="card__headline vacancy__link">`)
		qw422016.E().S(company.Name)
		qw422016.N().S(`</a>
                </figcaption>
              </figure>
              <div class="card__info">
                <figure class="card__figure">
                  <img
                    class="card__icon"
                    alt="card type icon"
                    width="16"
                    height="16"
                    src="/assets/images/pages/vacancy/building.svg"
                  />
                  <figcaption class="card__figcaption">`)
		qw422016.E().S(companyTypeName[company.Type])
		qw422016.N().S(`</figcaption>
                </figure>
                <figure class="card__figure">
                  <img
                    class="card__icon"
                    alt="card type icon"
                    width="16"
                    height="16"
                    src="/assets/images/pages/vacancy/company-type.svg"
                  />
                  <figcaption class="card__figcaption">`)
		qw422016.E().S(industryNames(company.Industries))
		qw422016.N().S(`</figcaption>
                </figure>
              </div>
              <p class="js-company-description card__text">`)
		qw422016.E().S(company.ShortDescription)
		qw422016.N().S(`</p>
              <div class="card__links">
                <ul class="card__links-group">
                  <li class="card__links-item">
                    <img
                      class="card__links-icon"
                      alt="linkedin icon"
                      width="20"
                      height="20"
                      src="/assets/images/pages/organizer/linkedin.svg"
                    />
                    <a href="https://www.linkedin.com/company/`)
		qw422016.E().S(company.LinkedInProfile.Alias)
		qw422016.N().S(`/" class="button-link card__links-link">LinkedIn</a>
                  </li>
                  <li class="card__links-item">
                    <a href="`)
		qw422016.E().S(linkedinConnectionsURL([]Company{company}, ukrainianUniversities))
		qw422016.N().S(`" class="button-link card__links-link">Connections             <img
                            class="checkbox__content-image"
                            alt="Flag of Ukraine with waves"
                            width="16"
                            height="16"
                            src="/assets/images/pages/online-new/ua_flag_with_waves.svg"
                    /></a>
                  </li>
                  <li class="card__links-item">
                    <a href="`)
		qw422016.E().S(linkedinConnectionsURL([]Company{company}, czechUniversities))
		qw422016.N().S(`" class="button-link card__links-link">Connections             <img
                            class="checkbox__content-image"
                            alt="Flag of Czechia"
                            width="16"
                            height="16"
                            src="/assets/images/pages/online-new/cz_flag.svg"
                    /></a>
                  </li>
                  <li class="card__links-item">
                    <a href="`)
		qw422016.E().S(linkedinConnectionsURL([]Company{company}, nil))
		qw422016.N().S(`" class="button-link card__links-link">Connections</a>
                  </li>
                  <li class="card__links-item">
                    <a href="`)
		qw422016.E().S(linkedinJobsURL([]Company{company}, string(organizerFeature.Organizer.LanguageTitleKeywords)))
		qw422016.N().S(`" class="button-link card__links-link">Jobs</a>
                  </li>
                </ul>
`)
		if company.GitHubProfile.Login == "" {
			qw422016.N().S(`                  <ul class="card__links-group">
                    <li class="card__links-item card__links-item--disabled">
                      <img
                        class="card__links-icon"
                        alt="github icon"
                        width="20"
                        height="20"
                        src="/assets/images/pages/organizer/github.svg"
                      />
                      <a href="`)
			qw422016.E().S(googleSearchGitHub(company.Name))
			qw422016.N().S(`" class="card__links-link card__links-link--google">
                        <img
                          class="card__links-icon card__links-icon--google"
                          alt="google icon"
                          width="20"
                          height="20"
                          src="/assets/images/pages/organizer/google.svg"
                        />
                      </a>
                      <span class="button-link card__links-link">GitHub</span>
                    </li>
                  </ul>
                `)
		} else {
			qw422016.N().S(`
                  <ul class="card__links-group">
                    <li class="card__links-item">
                      <img
                        class="card__links-icon"
                        alt="github icon"
                        width="20"
                        height="20"
                        src="/assets/images/pages/organizer/github.svg"
                      />
                      <a href="https://github.com/`)
			qw422016.E().S(company.GitHubProfile.Login)
			qw422016.N().S(`" class="button-link card__links-link">GitHub</a>
                    </li>
                    <li class="card__links-item">
                      `)
			qw422016.N().S(`
                      <a href="https://github.com/orgs/`)
			qw422016.E().S(company.GitHubProfile.Login)
			qw422016.N().S(`/repositories?q=lang:go" class="button-link card__links-link" title="`)
			qw422016.N().D(company.GitHubProfile.GoRepositoryCount)
			qw422016.N().S(` repositories">Repositories</a>
                    </li>
                  </ul>
`)
		}
		if company.GlassdoorProfile.OverviewURL == "" {
			qw422016.N().S(`                  <ul class="card__links-group">
                    <li class="card__links-item card__links-item--disabled">
                      <img
                        class="card__links-icon"
                        alt="glassdoor icon"
                        width="20"
                        height="20"
                        src="/assets/images/pages/organizer/glassdoor.svg"
                      />
                      <a href="`)
			qw422016.E().S(googleSearchGlassdoor(company.Name))
			qw422016.N().S(`" class="card__links-link card__links-link--google">
                        <img
                          class="card__links-icon card__links-icon--google"
                          alt="google icon"
                          width="20"
                          height="20"
                          src="/assets/images/pages/organizer/google.svg"
                        />
                      </a>
                      <span class="button-link card__links-link">Glassdoor</span>
                    </li>
                  </ul>
                `)
		} else {
			qw422016.N().S(`
                  <ul class="card__links-group">
                    <li class="card__links-item">
                      <img
                        class="card__links-icon"
                        alt="glassdoor icon"
                        width="20"
                        height="20"
                        src="/assets/images/pages/organizer/glassdoor.svg"
                      />
                      <a href="`)
			qw422016.E().S(company.GlassdoorProfile.OverviewURL)
			qw422016.N().S(`" class="button-link card__links-link">Glassdoor</a>
                    </li>
                    <li class="card__links-item">
                      <a href="`)
			qw422016.E().S(company.GlassdoorProfile.ReviewsURL)
			qw422016.N().S(`" class="button-link card__links-link">Reviews</a>
                    </li>
                  </ul>
`)
		}
		qw422016.N().S(`                <ul class="card__links-group card__links-group--unbordered">
                  <li class="card__links-item">
                    <img
                      class="card__links-icon"
                      alt="similarweb icon"
                      width="20"
                      height="20"
                      src="/assets/images/pages/organizer/similarweb.svg"
                    />
                    <a href="`)
		qw422016.E().S(similarwebURL(company.URL))
		qw422016.N().S(`" class="button-link card__links-link">SimilarWeb</a>
                  </li>
                  <li class="card__links-item">
                    <img
                      class="card__links-icon"
                      alt="whois icon"
                      width="20"
                      height="20"
                      src="/assets/images/pages/organizer/whois.svg"
                    />
                    <a href="`)
		qw422016.E().S(whoisURL(company.URL))
		qw422016.N().S(`" class="button-link card__links-link">Whois</a>
                  </li>

                  `)
		qw422016.N().S(`
                  <li class="card__links-item card__links-item--disabled">
                    <img
                      class="card__links-icon"
                      alt="xing icon"
                      width="20"
                      height="20"
                      src="/assets/images/pages/organizer/xing.svg"
                    />
                    <a href="`)
		qw422016.E().S(googleSearchXing(company.Name))
		qw422016.N().S(`" class="card__links-link card__links-link--google">
                      <img
                        class="card__links-icon card__links-icon--google"
                        alt="google icon"
                        width="20"
                        height="20"
                        src="/assets/images/pages/organizer/google.svg"
                      />
                    </a>
                    <span class="button-link card__links-link">XING</span>
                  </li>

`)
		if company.OttaProfileSlug == "" {
			qw422016.N().S(`                    <li class="card__links-item card__links-item--disabled">
                      <img
                        class="card__links-icon"
                        alt="otta icon"
                        width="20"
                        height="20"
                        src="/assets/images/pages/organizer/otta.svg"
                      />
                      <a href="`)
			qw422016.E().S(googleSearchOtta(company.Name))
			qw422016.N().S(`" class="card__links-link card__links-link--google">
                        <img
                          class="card__links-icon card__links-icon--google"
                          alt="google icon"
                          width="20"
                          height="20"
                          src="/assets/images/pages/organizer/google.svg"
                        />
                      </a>
                      <span class="button-link card__links-link">Otta</span>
                    </li>
                  `)
		} else {
			qw422016.N().S(`
                    <li class="card__links-item">
                      <img
                        class="card__links-icon"
                        alt="otta icon"
                        width="20"
                        height="20"
                        src="/assets/images/pages/organizer/otta.svg"
                      />
                      <a href="https://app.otta.com/companies/`)
			qw422016.E().S(company.OttaProfileSlug)
			qw422016.N().S(`" class="button-link card__links-link">Otta</a>
                    </li>
`)
		}
		qw422016.N().S(`                </ul>
              </div>
            </div>
            `)
	}
	qw422016.N().S(`
          </div>

          `)
	qw422016.N().S(`
        </div>
      </div>
    </div>
  </div>
</div>


</main>
`)
	streamorganizersFooter(qw422016)
	qw422016.N().S(`
<script src="/assets/js/organizers-companies-app.js?`)
	qw422016.N().D(appVersion)
	qw422016.N().S(`"></script>
</body>
</html>
`)
}

func WriteOrganizersCompanies(qq422016 qtio422016.Writer,
	organizerFeature OrganizerFeature,
	headerProfiles []SocialProviderUser,
	companies []Company,
	ukrainianUniversities []University,
	czechUniversities []University,
	userCompanyFavoriteMap map[int64]bool,
	authQueryParams string,
) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamOrganizersCompanies(qw422016, organizerFeature, headerProfiles, companies, ukrainianUniversities, czechUniversities, userCompanyFavoriteMap, authQueryParams)
	qt422016.ReleaseWriter(qw422016)
}

func OrganizersCompanies(
	organizerFeature OrganizerFeature,
	headerProfiles []SocialProviderUser,
	companies []Company,
	ukrainianUniversities []University,
	czechUniversities []University,
	userCompanyFavoriteMap map[int64]bool,
	authQueryParams string,
) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteOrganizersCompanies(qb422016, organizerFeature, headerProfiles, companies, ukrainianUniversities, czechUniversities, userCompanyFavoriteMap, authQueryParams)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
